package trigger

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// func (m *Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
//     m.mutex.Lock()
//     defer m.mutex.Unlock()

//     if m.running {
//         return errors.New("bot is already running")
//     }

//     discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
//     if err != nil {
//         return errors.New("error creating Discord session")
//     }

//     m.discord = discord
//     m.stop = make(chan struct{})
//     m.running = true

//     if err := m.discord.Open(); err != nil {
//         return errors.New("error opening connection")
//     }

//     log.Println("Bot started and running...")
//     return nil
// }

func (m *Model) Watch(ctx context.Context, userID string, actionNode workspace.ActionNodeModel) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()

    var userSession UserSession
    err := m.Collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&userSession)
    if err != nil && err != mongo.ErrNoDocuments {
        return fmt.Errorf("error finding user session: %v", err)
    }

    if userSession.Running {
        return errors.New("bot is already running for this user")
    }

    discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
    if err != nil {
        return errors.New("error creating Discord session")
    }

    stopChan := make(chan struct{})
    go m.runDiscordSession(discord, stopChan)

    // Update MongoDB with the new session data
    _, err = m.Collection.UpdateOne(
        ctx,
        bson.M{"user_id": userID},
        bson.M{"$set": UserSession{UserID: userID, Token: os.Getenv("BOT_TOKEN"), Running: true, StopChan: true}},
        options.Update().SetUpsert(true),
    )
    if err != nil {
        return fmt.Errorf("error storing user session: %v", err)
    }

    log.Printf("Bot started and running for user %s...\n", userID)
    return nil
}

func (m *Model) runDiscordSession(discord *discordgo.Session, stopChan chan struct{}) {
    defer discord.Close()

    err := discord.Open()
    if err != nil {
        log.Printf("error opening Discord session: %v", err)
        return
    }
    log.Println("Bot running...")

    <-stopChan // Wait until stop is triggered

    log.Println("Bot session stopped.")
}


func newMessage(s *discordgo.Session, m *discordgo.MessageCreate, stop chan struct{}) {

	log.Printf("Message received in GuildID: %s, ChannelID: %s\n", m.GuildID, m.ChannelID)
	log.Printf("Message from %s: %s\n", m.Author.Username, m.Content)
    select {
    case stop <- struct{}{}:
    default:
    }
}


// func (m *Model) Webhook(ctx context.Context) error {
//     m.mutex.Lock()
//     defer m.mutex.Unlock()

//     if !m.running || m.discord == nil {
//         return errors.New("bot is not running")
//     }

//     m.discord.AddHandler(func(s *discordgo.Session, msg *discordgo.MessageCreate) {
//         newMessage(s, msg, m.stop)
//     })

//     log.Println("Message handler added.")
//     return nil
// }

func (m *Model) Webhook(ctx context.Context, userID string) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()

    // Fetch the user's session information from MongoDB
    var userSession UserSession
    err := m.Collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&userSession)
    if err != nil {
        return fmt.Errorf("error finding user session: %v", err)
    }

    // Check if the bot is running for the user
    if !userSession.Running {
        return errors.New("bot is not running for this user")
    }

    // Create a new Discord session for this user
    discord, err := discordgo.New("Bot " + userSession.Token)
    if err != nil {
        return errors.New("error creating Discord session")
    }

    // Add a message handler to respond to messages
    discord.AddHandler(func(s *discordgo.Session, msg *discordgo.MessageCreate) {
        // Pass the stop channel from MongoDB if necessary
        newMessage(s, msg, make(chan struct{})) // MongoDB doesn’t handle channels, so instantiate a temporary channel for handling
    })

    // Open the Discord session to start receiving messages
    if err := discord.Open(); err != nil {
        return errors.New("error opening connection")
    }

    log.Printf("Message handler added for user %s.\n", userID)
    return nil
}



// func (m *Model) Stop(ctx context.Context) error {
//     m.mutex.Lock()
//     defer m.mutex.Unlock()

//     if !m.running || m.discord == nil {
//         return errors.New("bot is not running")
//     }

//     m.discord.Close()

//     m.running = false
//     m.discord = nil

//     m.stop = nil

//     log.Println("Bot stopped.")
//     return nil
// }

func (m *Model) Stop(ctx context.Context, userID string) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()

    var userSession UserSession
    err := m.Collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&userSession)
    if err != nil {
        return fmt.Errorf("error finding user session: %v", err)
    }

    if !userSession.Running {
        return errors.New("bot is not running for this user")
    }

    // Update the session state in MongoDB
    _, err = m.Collection.UpdateOne(
        ctx,
        bson.M{"user_id": userID},
        bson.M{"$set": bson.M{"running": false, "stop_chan": false}},
    )
    if err != nil {
        return fmt.Errorf("error updating user session: %v", err)
    }

    log.Printf("Bot stopped for user %s.\n", userID)
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