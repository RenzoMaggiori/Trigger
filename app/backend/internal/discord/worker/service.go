package worker

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"trigger.com/trigger/internal/sync"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func (m *Model) Me(token string) (*Me, error) {
	user, _, err := user.GetUserByAccesstokenRequest(token)
	if err != nil {
		return nil, err
	}

	sync, _, err := sync.GetSyncAccessTokenRequest(token, user.Id.Hex(), "discord")
	log.Println("sync_access_token", sync.AccessToken)
	if err != nil {
		return nil, err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			userEndpoint,
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", sync.AccessToken),
			},
		),
	)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDiscordMe, err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %v", errors.ErrDiscordMe, res.StatusCode)
	}

	discord_me, err := decode.Json[DiscordMe](res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDecodeData, err)
	}

	log.Println("id", discord_me.Id)
	log.Println("username", discord_me.Username)
	log.Println("email", discord_me.Email)

	me := Me{
		DiscordId:       discord_me.Id,
		Username: discord_me.Username,
		Email:    discord_me.Email,
	}

	return &me, nil
}

func (m *Model) GuildChannels(guildID string) ([]Channel, error) {
	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrCreateDiscordGoSession, err)
	}

	defer discord.Close()

	channels, err := discord.GuildChannels(guildID)

	if err != nil {
		return nil, err
	}

	var response []Channel
	for _, ch := range channels {
		if ch.Type != discordgo.ChannelTypeGuildText {
			continue
		}
		response = append(response, Channel{
			Id:   ch.ID,
			Name: ch.Name,
		})
	}

	return response, nil
}

func (m *Model) AddSession(session *AddDiscordSession) error {
	ctx := context.TODO()
	newSync := DiscordSessionModel{
		UserId:  session.UserId,
		DiscordId: session.DiscordId,
		GuildId: session.GuildId,
		Token:   os.Getenv("BOT_TOKEN"),
		Running: false,
		Stop:    true,
	}

	_, err := m.Collection.InsertOne(ctx, newSync)
	if err != nil {
		return errors.ErrAddDiscordSession
	}

	log.Printf("Discord session created for user %s...\n", session.UserId)

	return nil
}

func (m *Model) UpdateSession(userId string, session *UpdateDiscordSession) error {
	ctx := context.TODO()
	filter := bson.M{"user_id": userId}
	update := bson.M{"$set": bson.M{"running": session.Running, "stop": session.Stop}}

	_, err := m.Collection.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return errors.ErrUpdateDiscordSession
	}

	return nil
}

func  (m *Model) GetSession(token string) (*DiscordSessionModel, error) {
	user, _, err := user.GetUserByAccesstokenRequest(token)
	if err != nil {
		return nil, err
	}

	var discordSession DiscordSessionModel
	err = m.Collection.FindOne(context.TODO(), bson.M{"user_id": user.Id.Hex()}).Decode(&discordSession)
	if err != nil {
		return nil, errors.ErrDiscordUserSessionNotFound
	}

	return &discordSession, nil
}

func (m *Model) DeleteSession(userId string) error {
	ctx := context.TODO()
	filter := bson.M{"user_id": userId}

	_, err := m.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.ErrDeleteDiscordSession
	}

	log.Printf("Discord session deleted for user %s...\n", userId)

	return nil
}

// func (m *Model) Watch(ctx context.Context, actionNode sync.ActionNodeModel) error {
// 	m.mutex.Lock()
// 	defer m.mutex.Unlock()

// 	user, err := getCurrUser(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	discordSession, err := m.getSession(user.AccessToken)
// 	if err != nil {
// 		return err
// 	}

// 	if discordSession.Running {
// 		return errors.ErrBotAlreadyRunning
// 	}

// 	discord, err := trigger.CreateDiscordSession()
// 	if err != nil {
// 		return err
// 	}

// 	stop := make(chan struct{})
// 	go m.runDiscordSession(discord, stop)

// 	_, err = m.Collection.UpdateOne(
// 		ctx,
// 		bson.M{"user_id": user.Id},
// 		bson.M{"$set": DiscordSessionModel{UserId: user.Id.Hex(), Token: os.Getenv("BOT_TOKEN"), Running: true, Stop: true}},
// 	)
// 	if err != nil {
// 		return fmt.Errorf("error storing user session: %v", err)
// 	}

// 	log.Printf("Bot started and running for user %s...\n", user.Id.Hex())

// 	return nil
// }