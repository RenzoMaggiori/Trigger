package trigger

import (
	"context"
	"errors"
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

	discord.AddHandler(newMessage)

	err = discord.Open()
	defer discord.Close()
	if err != nil {
		return errors.New("error opening connection")
	}

	log.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	return nil
}

func newMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	 // Identify which server the message is coming from
	log.Printf("Message received in GuildID: %s, ChannelID: %s\n", m.GuildID, m.ChannelID)

	log.Printf("Message from %s: %s\n", m.Author.Username, m.Content)

	// Respond to user
	// s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hello, %s! %s", m.Author.Username, "userInfo"))
	// s.ChannelMessageSendComplex(channelID, message)

}

func (m Model) Webhook(ctx context.Context) error {
	return nil
}

func (m Model) Stop(ctx context.Context) error {
	return nil
}