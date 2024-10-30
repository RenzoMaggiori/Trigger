package trigger

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/discord/worker"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/middleware"
)

func (m *Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()


    discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
    if err != nil {
        return errors.ErrCreateDiscordGoSession
    }

    m.bot = discord


    accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
    if !ok {
        return errors.ErrAccessTokenCtx
    }

    user, _, err := user.GetUserByAccesstokenRequest(accessToken)
    if err != nil {
        return err
    }

    me, err := worker.GetMeReq(accessToken)
    if err != nil {
        return err
    }
    currSession, err := worker.GetCurrDiscordSessionReq(accessToken)
    if err != nil {
        newSession := worker.AddDiscordSessionModel{
            UserId:  user.Id.Hex(),
            DiscordId: me.DiscordId,
            GuildId: "1300628531725991976",
        }

        err = worker.AddDiscordSessionReq(accessToken, newSession)
        if err != nil {
            return err
        }

        currSession, err = worker.GetCurrDiscordSessionReq(accessToken)
        if err != nil {
            return err
        }
    }

    if currSession.Running {
        return errors.ErrBotAlreadyRunning
    }


    if err := m.bot.Open(); err != nil {
        return errors.ErrOpeningDiscordConnection
    }

    err = worker.UpdateDiscordSessionReq(accessToken, user.Id.Hex(), worker.UpdateDiscordSessionModel{
        Running: true,
        Stop:    false,
    })
    if err != nil {
        log.Println("**** error updating session ****")
        return err
    }

    log.Println("Bot started and running...")
    return nil
}

func (m *Model) newMessage(s *discordgo.Session, msg *discordgo.MessageCreate, accessToken string) {
    log.Println("NEW MESSAGE RECEIVED")

    if msg.Author.ID == s.State.User.ID {
        log.Println("Message from bot itself, ignoring.")
        return
    }

    m.mutex.Lock()
    defer m.mutex.Unlock()

    currSession, err := worker.GetCurrDiscordSessionReq(accessToken)
    if err != nil {
        log.Printf("Error getting current session: %v", err)
        return
    }

    if !currSession.Running {
        log.Println("Current session is not running, ignoring message.")
        return
    }

    if msg.GuildID != currSession.GuildId {
        log.Printf("Message from different guild: %s, expected: %s", msg.GuildID, currSession.GuildId)
        return
    }

    log.Printf("Message received in GuildID: %s, ChannelID: %s\n", msg.GuildID, msg.ChannelID)
    log.Printf("Message from %s: %s\n", msg.Author.Username, msg.Content)
}


func (m *Model) Webhook(ctx context.Context) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()

    accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
    if !ok {
        return errors.ErrAccessTokenCtx
    }

    currSession, err := worker.GetCurrDiscordSessionReq(accessToken)
    if err != nil {
        return err
    }

    if !currSession.Running {
        return errors.ErrBotNotRunning
    }

    if m.bot == nil {
        return fmt.Errorf("discord session is nil")
    }

    m.bot.AddHandler(func(s *discordgo.Session, msg *discordgo.MessageCreate) {
        m.newMessage(s, msg, accessToken)
    })

    log.Println("Message handler added.")
    return nil
}

func (m *Model) Stop(ctx context.Context) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()

    accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
    if !ok {
        return errors.ErrAccessTokenCtx
    }

    user, _, err := user.GetUserByAccesstokenRequest(accessToken)
    if err != nil {
        return err
    }

    currSession, err := worker.GetCurrDiscordSessionReq(accessToken)
    if err != nil {
        return err
    }

    if !currSession.Running {
        return errors.ErrBotNotRunning
    }

    err = worker.UpdateDiscordSessionReq(accessToken, user.Id.Hex(), worker.UpdateDiscordSessionModel{
        Running: false,
        Stop:    true,
    })

    if err != nil {
        return err
    }

    log.Printf("Bot stopped for user %s.\n", user.Id.Hex())
    return nil
}
