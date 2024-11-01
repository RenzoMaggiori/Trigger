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
	errAction             error = errors.New("discord action not found")
	errWebhookBadStatus   error = errors.New("webhook returned a bad status")
)

func (m *Model) New(ctx context.Context) *cron.Cron {
	m.InitBot()
	log.Println("-------------------------\n\nONLY ONCE\n\n-------------------------")
	c := cron.New()
	err := c.AddFunc("0 */1 * * * *", func() {
	log.Println("-------------------------\n\n[ **DISCORD JOB RUNNING **]\n\n-------------------------")

		if err := m.Start(ctx); err != nil {
			log.Println("error starting discord job: ", err)
		}
		log.Println("-------------------------\n\n[ ** DISCORD JOB RUNNING **]\n\n-------------------------")
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

	// log.Printf("----------------------------------------------------------\n\n")
	// log.Printf("					WORKSPACES\n\n")
	// prettyJSON, err := json.MarshalIndent(workspaces, "", "    ")
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(string(prettyJSON))
	// log.Printf("\n\n----------------------------------------------------------\n\n")


	var wg sync.WaitGroup

	// for (workspaces) == for each user_id
	for _, w := range workspaces {
		wg.Add(1)

		log.Printf("----------------------------------------------------------\n\n")
		log.Printf("					WORKSPACE\n\n")
		prettyJSON, err := json.MarshalIndent(w, "", "    ")
		if err != nil {
			return err
		}
		fmt.Println(string(prettyJSON))
		log.Printf("\n\n----------------------------------------------------------\n\n")

		go func(w workspace.WorkspaceModel, a action.ActionModel) {
			defer wg.Done()
			err := m.handleUserNodes(w, a)
			if err != nil {
				log.Printf("Error processing user %s: %v", w.UserId.Hex(), err)
			}
		}(w, *discordAction)
	}
	wg.Wait()
	return nil
}

func (m *Model) handleUserNodes(w workspace.WorkspaceModel, a action.ActionModel) error {
	userId := w.UserId.Hex()
	userTokens, err := getUserAccessToken(userId)
	if err != nil {
		return err
	}
	log.Printf("-------------\n\nUSER TOKENS\n\n")
	log.Println("sync ", userTokens.sync.AccessToken)
	log.Println("session", userTokens.session.AccessToken)
	log.Println("S\n\n--------------")


	log.Printf("-------------\n\nUSER ID: %s \n\n--------------\n\n", userId)
	actionId := a.Id.Hex()
	log.Printf("-------------\n\nACTION ID: %s \n\n--------------\n\n", actionId)

	if len(w.Nodes) == 0 {
		return nil
	}

	//* stores the last one
	// var input map[string]string
	// for _, n := range w.Nodes {
	// 	if n.ActionId.Hex() != actionId {
	// 		continue
	// 	}
	// 	input = n.Input
	// }

	//* stores all that accomplish the condition
	var inputs []map[string]string
	for _, n := range w.Nodes {
		if n.ActionId.Hex() == actionId {
			inputs = append(inputs, n.Input)
		}
	}

	if inputs == nil {
		return nil
	}

	for _, input := range inputs {
		channel_id, ok := input["channel_id"]
		if !ok {
			return nil
		}

		discord_me, err := m.GetMe(userTokens.sync.AccessToken)
		if err != nil {
			return err
		}

		existingSession, _ := m.GetSessionByUserId(userId)
		if existingSession != nil {
			if existingSession.ChannelId == channel_id {
				log.Println("Session already exists for this channel")
				continue
			}
		}

		newSession := &DiscordSessionModel{
			UserId:    userId,
			ChannelId: channel_id,
			ActionId:  actionId,
			Token:     userTokens.session.AccessToken,
			DiscordData: discord_me,
		}
		err = m.AddSession(newSession)
		if err != nil {
			log.Printf("Error adding session [%s]: %v", newSession.ChannelId, err)
			// return err
		}
	}
	// channel_id, ok := input["channel_id"]
	// if !ok {
	// 	return nil
	// }

	// discord_me, err := m.GetMe(userTokens.sync.AccessToken)
	// if err != nil {
	// 	return err
	// }

	// newSession := &DiscordSessionModel{
	// 	UserId:    userId,
	// 	DiscordData: discord_me,
	// 	ChannelId: channel_id,
	// 	ActionId:  actionId,
	// 	Token:     userTokens.session.AccessToken,
	// }
	// err = m.AddSession(newSession)
	// if err != nil {
	// 	return err
	// }
	return nil
}

// * SESSION TOKEN // (discord) SYNC TOKEN
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
	syncModel, _, err := userSync.GetSyncAccessTokenRequest(userTokens.session.AccessToken, userId, "discord")
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
		if a.Action != "watch_channel_message" {
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

	log.Println("THERE IS A NEW MESSAGE AVAILABLE --> FETCHING DISCORD WEBHOOK...")
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
	log.Println("-------------------------\n\nNEW MESSAGE RECEIVED\n\n-------------------------")

	if msg.Author.ID == s.State.User.ID {
		log.Println("Message from bot itself, ignoring.")
		return
	}

	// m.mutex.Lock()
	// defer m.mutex.Unlock()

	discord_sessions, err := m.GetAllDiscordSessions()
	if err != nil {
		log.Printf("Error getting all discord sessions: %v", err)
		return
	}
	log.Printf("-------------------------\n\nDISCORD SESSIONS\n\n")
	log.Println(discord_sessions)
	log.Printf("-------------------------\n\n")


	for _, ds := range discord_sessions {
		if ds.ChannelId == msg.ChannelID {
			log.Printf("-------------------------\n\n")
			log.Printf("Message received in  ChannelID: %s\n", msg.ChannelID)
			log.Printf("Message from %s: %s\n", msg.Author.Username, msg.Content)
			log.Printf("-------------------------\n\n")

			err := m.FetchDiscordWebhook(ds.Token, trigger.ActionBody{
				Type: "watch_message",
				Data: trigger.MsgInfo{
					Content: msg.Content,
					AuthoId:  msg.Author.ID,
					AuthoUsername: msg.Author.Username,
				},
			})
			if err != nil {
				return
			}
		}
	}

}

func (m *Model) InitBot() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var err error
	m.bot, err = discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		return myErrors.ErrCreateDiscordGoSession
	}

	if err := m.bot.Open(); err != nil {
		return myErrors.ErrOpeningDiscordConnection
	}
	// defer m.bot.Close()

	// m.bot.AddHandler(func(s *discordgo.Session, msg *discordgo.MessageCreate) {
	// 	log.Println("New message received")
	// 	m.HandleNewMessage(s, msg)
	// })

	m.bot.AddHandler(m.HandleNewMessage)

	log.Println("-------------------------\n\nBot started and running...\n\n-------------------------")
	return nil
}
