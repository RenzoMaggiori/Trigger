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

func createDiscordSession() (*discordgo.Session, error) {
    discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
    if err != nil {
        return nil, errors.New("error creating Discord session")
    }
    return discord, nil
}

func addMessageHandler(discord *discordgo.Session, stop chan struct{}) {
    discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
        newMessage(s, m, stop)
    })
}

func watchSession(ctx context.Context, discord *discordgo.Session, stop chan struct{}) error {
    if err := discord.Open(); err != nil {
        return errors.New("error opening connection")
    }
    defer discord.Close()

    log.Println("Bot running...")

    select {
    case <-stop:
    case <-ctx.Done():
    }

    return nil
}

func (m *Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()

    if m.running {
        return errors.New("bot is already running")
    }

    discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
    if err != nil {
        return errors.New("error creating Discord session")
    }

    m.discord = discord
    m.stop = make(chan struct{})
    m.running = true

    if err := m.discord.Open(); err != nil {
        return errors.New("error opening connection")
    }

    log.Println("Bot started and running...")
    return nil
}

func newMessage(s *discordgo.Session, m *discordgo.MessageCreate, stop chan struct{}) {

	log.Printf("Message received in GuildID: %s, ChannelID: %s\n", m.GuildID, m.ChannelID)
	log.Printf("Message from %s: %s\n", m.Author.Username, m.Content)
    select {
    case stop <- struct{}{}:
    default:
    }
}


func (m *Model) Webhook(ctx context.Context) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()

    if !m.running || m.discord == nil {
        return errors.New("bot is not running")
    }

    m.discord.AddHandler(func(s *discordgo.Session, msg *discordgo.MessageCreate) {
        newMessage(s, msg, m.stop)
    })

    log.Println("Message handler added.")
    return nil
}

func (m *Model) Stop(ctx context.Context) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()

    if !m.running || m.discord == nil {
        return errors.New("bot is not running")
    }

    m.discord.Close()

    m.running = false
    m.discord = nil

    m.stop = nil

    log.Println("Bot stopped.")
    return nil
}


func (m *Model) Guilds() error {
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

func (m *Model) GuildChannels(guildID string) error {
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