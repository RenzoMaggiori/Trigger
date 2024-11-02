package worker

import (
	"bytes"
	// "context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	// "sync"

	"github.com/bwmarrin/discordgo"
	"trigger.com/trigger/internal/discord/trigger"
	myErrors "trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

var (
	errWebhookBadStatus   error = errors.New("webhook returned a bad status")
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
			log.Printf("Message received in  ChannelID: %s\n", msg.ChannelID)
			log.Printf("Message from %s: %s\n", msg.Author.Username, msg.Content)
			log.Printf("-------------------------\n\n")

			err := m.FetchDiscordWebhook(ds.Token, trigger.ActionBody{
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
