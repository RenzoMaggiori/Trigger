package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/discord/trigger"
	"trigger.com/trigger/internal/session"
	userSync "trigger.com/trigger/internal/sync"
	myErrors "trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/mongodb"
)

var (
	errCollectionNotFound error = errors.New("could not find discord collection")
	errSessionNotFound    error = errors.New("could not find user session")
	errSyncModelNull      error = errors.New("the sync models type is null")
	errAction      error = errors.New("discord action not found")
	errWebhookBadStatus   error = errors.New("webhook returned a bad status")
)

func (m *Model) New(ctx context.Context) *cron.Cron {
	m.InitBot()
	c := cron.New()
	err := c.AddFunc("0 */1 * * * *", func() {
		log.Println("discord job running...")
		if err := m.Start(ctx); err != nil {
			log.Println(err)
		}
		log.Println("discord job ended")
	})
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func (m *Model) Start(ctx context.Context) error {
	var ok bool
	m.Collection, ok = ctx.Value(mongodb.CtxKey).(*mongo.Collection)
	if !ok || m.Collection == nil {
		return errCollectionNotFound
	}

	discordAction, err := getDiscordAction()
	if err != nil {
		return err
	}

	workspaces, _, err := workspace.GetByActionIdRequest(
		os.Getenv("ADMIN_TOKEN"),
		discordAction.Id.Hex(),
	)
	if err != nil {
		return err
	}

	log.Printf("%v\n", workspaces)

	var wg sync.WaitGroup

	// for (workspaces) == for each user_id
	for _, w := range workspaces {
		wg.Add(1)
		go func(w workspace.WorkspaceModel, a action.ActionModel) {
			defer wg.Done()
			err := m.newMessageRecieved(w, a)
			if err != nil {
				log.Printf("Error processing user %s: %v", w.UserId.Hex(), err)
			}
		}(w, *discordAction)
	}
	wg.Wait()
	return nil
}

func (m *Model) newMessageRecieved(w workspace.WorkspaceModel, a action.ActionModel) error {
	userId := w.UserId.Hex()

	actionId := a.Id.Hex()

	if len(w.Nodes) == 0 {
		return nil
	}

	var input map[string]string
	for _, n := range w.Nodes {
		if n.ActionId.Hex() != actionId {
			continue
		}
		input = n.Input
	}
	if input == nil {
		return nil
	}

	channel_id, ok := input["channel_id"]
	if !ok {
		return nil
	}
	guild_id, ok := input["guild_id"]
	if !ok {
		return nil
	}

	userTokens, err := getUserAccessToken(userId)
	if err != nil {
		return err
	}

	discord_me, err := m.GetMe(userTokens.sync.AccessToken)
	if err != nil {
		return err
	}

	newSession := &DiscordSessionModel{
		UserId: userId,
		DiscordId: discord_me.DiscordId,
		GuildId: guild_id,
		ChannelId: channel_id,
		ActionId: actionId,
		Token: userTokens.session.AccessToken,
	}

	err = m.AddSession(newSession)
	if err != nil {
		return err
	}

	return nil
}

//* SESSION TOKEN // SYNC TOKEN (discord)
func getUserAccessToken(userId string) (*UserTokens, error) {
	session, _, err := session.GetSessionByUserIdRequest(os.Getenv("ADMIN_TOKEN"), userId)
	if err != nil {
		return nil, err
	}
	if len(session) == 0 {
		return nil, errSessionNotFound
	}

	userTokens := UserTokens{
		session: session[0],
	}
	syncModel, _, err := userSync.GetSyncAccessTokenRequest(userTokens.session.AccessToken, userId, "spotify")
	if err != nil {
		return nil, err
	}
	if syncModel == nil {
		return nil, errSyncModelNull
	}

	userTokens.sync = *syncModel
	return &userTokens, nil
}

func getDiscordAction() (*action.ActionModel, error) {
	actions, _, err := action.GetByProviderRequest(
		os.Getenv("ADMIN_TOKEN"),
		"discord",
	)
	if err != nil {
		return nil, err
	}

	for _, a := range actions {
		if a.Type != "trigger" {
			continue
		}
		if a.Action != "watch_message" {
			continue
		}
		return &a, nil
	}
	return nil, errAction
}

func (m *Model) FetchDiscordWebhook(accessToken string, data trigger.ActionBody) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/discord/trigger/webhook", os.Getenv("DISCORD_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errWebhookBadStatus
	}
	return nil
}

func (m *Model) HandleNewMessage(s *discordgo.Session, msg *discordgo.MessageCreate) {
    log.Println("NEW MESSAGE RECEIVED")

    if msg.Author.ID == s.State.User.ID {
        log.Println("Message from bot itself, ignoring.")
        return
    }

    m.mutex.Lock()
    defer m.mutex.Unlock()

	discord_sessions, err := m.GetAllDiscordSessions()
	if err != nil {
		log.Printf("Error getting all discord sessions: %v", err)
		return
	}

	for _, ds := range discord_sessions {
		if ds.ChannelId == msg.ChannelID {
			log.Printf("Message received in GuildID: %s, ChannelID: %s\n", msg.GuildID, msg.ChannelID)
			log.Printf("Message from %s: %s\n", msg.Author.Username, msg.Content)

			err := m.FetchDiscordWebhook(ds.Token, trigger.ActionBody{
				Type: "watch_message",
				Data: trigger.MsgInfo{
					Author: msg.Author.Username,
					Content: msg.Content,
				},
			})
			if err != nil {
				return
			}
		}
	}

}

func (m *Model)InitBot() error {
    m.mutex.Lock()
    defer m.mutex.Unlock()

    bot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
    if err != nil {
        return myErrors.ErrCreateDiscordGoSession
    }

    if err := bot.Open(); err != nil {
        return myErrors.ErrOpeningDiscordConnection
    }

	defer bot.Close()

	bot.AddHandler(func(s *discordgo.Session, msg *discordgo.MessageCreate) {
        m.HandleNewMessage(s, msg)
    })

    log.Println("Bot started and running...")
    return nil
}

// func GetMembers(s *discordgo.Session, guildID string) ([]*discordgo.Member, error) {
// 	var after string
// 	var members []*discordgo.Member
// 	for {
// 		m, err := s.GuildMembers(guildID, after, 1000)
// 		if err != nil {
// 			log.Fatalf("error fetching members: %v", err)
// 		}

// 		members = append(members, m...)

// 		if len(m) < 1000 {
// 			break
// 		}

// 		after = m[len(m)-1].User.ID
// 	}

// 	// Print member usernames
// 	for _, member := range members {
// 		fmt.Printf("User: %s#%s\n", member.User.ID, member.User.Username)
// 	}

// 	return members, nil
// }