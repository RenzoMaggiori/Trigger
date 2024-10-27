package trigger

import (
	"context"
	"errors"
	"fmt"
	"syscall"

	// "fmt"

	// "fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"trigger.com/trigger/internal/action/workspace"
)

func (m Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		return errors.New("error creating Discord session")
	}
	defer discord.Close()


	// Create a channel to signal shutdown when a message is received
	stop := make(chan struct{})

	// Pass the stop to newMessage to signal shutdown on message receipt
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		newMessage(s, m, stop)
	})

	if err := discord.Open(); err != nil {
		return errors.New("error opening connection")
	}

	log.Println("Bot running....")

	// Listen for OS interrupt signals and the stopChan
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stop:
		log.Println("Message received; stopping bot.")
	case <-signalChan:
		log.Println("Interrupt signal received; stopping bot.")
	case <-ctx.Done():
		log.Println("Context canceled; stopping bot.")
	}

	return nil
}

func newMessage(s *discordgo.Session, m *discordgo.MessageCreate, stopChan chan struct{}) {
	// if m.Author.ID == s.State.User.ID {
	// 	return
	// }
	log.Printf("Message received in GuildID: %s, ChannelID: %s\n", m.GuildID, m.ChannelID)

	// Check for a non-bot message to proceed
	if m.Author.Bot {
		return
	}

	log.Printf("Message from %s: %s\n", m.Author.Username, m.Content)

	// Respond to user
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hello, %s! %s", m.Author.Username, "what's up!"))
	// s.ChannelMessageSendComplex(channelID, message)

	// Signal the main process to shut down
	stopChan <- struct{}{}

}

func (m Model) Webhook(ctx context.Context) error {
	return nil
}

func (m Model) Stop(ctx context.Context) error {
	return nil
}