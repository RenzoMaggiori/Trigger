package trigger

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"trigger.com/trigger/internal/action/workspace"
)

func (m Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		return errors.New("error creating Discord session")
	}
	defer discord.Close()

	stop := make(chan struct{})

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		newMessage(s, m, stop)
	})

	if err := discord.Open(); err != nil {
		return errors.New("error opening connection")
	}

	log.Println("Bot running....")

	select {
	case <-stop:
	case <-ctx.Done():
	}

	return nil
}

func newMessage(s *discordgo.Session, m *discordgo.MessageCreate, stopChan chan struct{}) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	log.Printf("Message received in GuildID: %s, ChannelID: %s\n", m.GuildID, m.ChannelID)
	log.Printf("Message from %s: %s\n", m.Author.Username, m.Content)

	// s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hello, %s! %s", m.Author.Username, "what's up!"))
	// s.ChannelMessageSendComplex(channelID, message)

	stopChan <- struct{}{}

}

func (m Model) Webhook(ctx context.Context) error {
	return nil
}

func (m Model) Stop(ctx context.Context) error {
	return nil
}


func (m Model) Guilds() error {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		return errors.New("error creating Discord session")
	}
	defer discord.Close()


	discord.AddHandler(func(s *discordgo.Session, event *discordgo.GuildCreate) {
		fmt.Printf("Guild ID: %s, Name: %s\n", event.Guild.ID, event.Guild.Name)
	})

	if err := discord.Open(); err != nil {
		return errors.New("error opening connection")
	}

	return nil
}

func (m Model) GuildChannels(guildID string) error {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		return errors.New("error creating Discord session")
	}

	defer discord.Close()

	channels, err := discord.GuildChannels(guildID)

	if err != nil {
		return err
	}

	for _, channel := range channels {
		if channel.Type != discordgo.ChannelTypeGuildText {
			continue
		}
		fmt.Printf("Channel ID: %s, Name: %s\n", channel.ID, channel.Name)
	}

	return nil

}