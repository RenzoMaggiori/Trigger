package worker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"trigger.com/trigger/internal/discord/trigger"
	"trigger.com/trigger/internal/session"
	myErrors "trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
	userSync "trigger.com/trigger/internal/sync"
)

var (
	errWebhookBadStatus   error = errors.New("webhook returned a bad status")
	errSyncModelNull      error = errors.New("the sync models type is null")
	errSessionNotFound    error = errors.New("could not find user session")
)

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

	m.mutex.Lock()
	defer m.mutex.Unlock()

	discord_sessions, err := trigger.GetAllDiscordSessions()
	if err != nil {
		log.Printf("Error getting all discord sessions: %v", err)
		return
	}

	for _, ds := range discord_sessions {
		if ds.ChannelId == msg.ChannelID {
			log.Printf("-------------------------\n\n")
			log.Printf("Message received in ChannelID: %s\n", msg.ChannelID)
			log.Printf("Message from [%s]: %s\n", msg.Author.Username, msg.Content)
			log.Printf("-------------------------\n\n")

			userTokens, err := getUserAccessToken(ds.UserId)

			err = m.FetchDiscordWebhook(userTokens.session.AccessToken, trigger.ActionBody{
				Type: "watch_channel_message",
				Data: trigger.MsgInfo{
					Content: msg.Content,
					AuthorId:  msg.Author.ID,
					AuthorUsername: msg.Author.Username,
					NodeId: ds.NodeId,
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

	m.bot.AddHandler(m.HandleNewMessage)

	log.Println("-------------------------\n\nBot started and running...\n\n-------------------------")
	return nil
}

func getUserAccessToken(userId string) (*UserTokens, error) {
	session, _, err := session.GetSessionByUserIdRequest(os.Getenv("ADMIN_TOKEN"), userId)
	if err != nil {
		return nil, err
	}
	if session == nil || len(session) == 0 {
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