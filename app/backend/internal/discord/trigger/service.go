package trigger

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/discord/worker"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/middleware"
)

func createDiscordSession() (*discordgo.Session, error) {
    discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
    if err != nil {
        return nil, errors.ErrCreateDiscordGoSession
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
        return errors.ErrOpeningDiscordConnection
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

func (m *Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()

    user, err := user.GetCurrUserByCtxRequest(ctx)
    if err != nil {
        return err
    }

    // var discordSession worker.DiscordSessionModel
    discordSession worker.Get

    if discordSession.Running {
        return errors.ErrBotAlreadyRunning
    }

    discord, err := createDiscordSession()
    if err != nil {
        return err
    }

    stop := make(chan struct{})
    go m.runDiscordSession(discord, stop)

    _, err = m.Collection.UpdateOne(
        ctx,
        bson.M{"user_id": user.Id},
        bson.M{"$set": worker.DiscordSessionModel{UserID: user.Id.Hex(), Token: os.Getenv("BOT_TOKEN"), Running: true, Stop: true}},
        options.Update().SetUpsert(true),
    )
    if err != nil {
        return fmt.Errorf("error storing user session: %v", err)
    }

    log.Printf("Bot started and running for user %s...\n", user.Id.Hex())

    return nil
}

func (m *Model) runDiscordSession(discord *discordgo.Session, stop chan struct{}) {
    defer discord.Close()
    err := discord.Open()
    if err != nil {
        log.Printf("error opening Discord session: %v", err)
        return
    }
    log.Println("Bot running...")

    <-stop

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

func (m *Model) Webhook(ctx context.Context) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()

    user, err := user.GetCurrUserByCtxRequest(ctx)
    if err != nil {
        return err
    }

    var discordSession worker.DiscordSessionModel
    err = m.Collection.FindOne(ctx, bson.M{"user_id": user.Id}).Decode(&discordSession)
    if err != nil {
        return fmt.Errorf("error finding user session: %v", err)
    }

    if !discordSession.Running {
        return errors.ErrBotNotRunning
    }

    discord, err := discordgo.New("Bot " + discordSession.Token)
    if err != nil {
        return errors.ErrCreateDiscordGoSession
    }

    discord.AddHandler(func(s *discordgo.Session, msg *discordgo.MessageCreate) {
        newMessage(s, msg, make(chan struct{}))
    })

    if err := discord.Open(); err != nil {
        return errors.ErrOpeningDiscordConnection
    }

    log.Printf("Message handler added for user %s.\n", user.Id.Hex())
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

func (m *Model) Stop(ctx context.Context) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()

    user, err := user.GetCurrUserByCtxRequest(ctx)
    if err != nil {
        return err
    }

    var discordSession worker.DiscordSessionModel
    err = m.Collection.FindOne(ctx, bson.M{"user_id": user.Id}).Decode(&discordSession)
    if err != nil {
        return fmt.Errorf("error finding user session: %v", err)
    }

    if !discordSession.Running {
        return errors.ErrBotNotRunning
    }

    _, err = m.Collection.UpdateOne(
        ctx,
        bson.M{"user_id": user.Id},
        bson.M{"$set": bson.M{"running": false, "stop_chan": false}},
    )
    if err != nil {
        return fmt.Errorf("error updating user session: %v", err)
    }

    log.Printf("Bot stopped for user %s.\n", user.Id.Hex())
    return nil
}
