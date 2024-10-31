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

func (m *Model) GetMe(token string) (*DiscordMe, error) {
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

	discord_struct, err := decode.Json[DiscordStruct](res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDecodeData, err)
	}

	log.Println("id", discord_struct.Id)
	log.Println("username", discord_struct.Username)
	log.Println("email", discord_struct.Email)

	discord_me := DiscordMe{
		DiscordId:       discord_struct.Id,
		Username: discord_struct.Username,
		Email:    discord_struct.Email,
	}

	return &discord_me, nil
}

//* GET CHANNELS FROM A GUILD
func (m *Model) GetGuildChannels(guildID string) ([]Channel, error) {
	bot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrCreateDiscordGoSession, err)
	}

	if err := bot.Open(); err != nil {
		return nil, errors.ErrOpeningDiscordConnection
    }
	defer bot.Close()

	channels, err := bot.GuildChannels(guildID)
	if err != nil {
		return nil, err
	}

	var response []Channel
	for _, ch := range channels {
		if ch.Type != discordgo.ChannelTypeGuildText {
			continue
		}
		response = append(response, Channel{Id: ch.ID, Name: ch.Name})
	}

	return response, nil
}

// * GET GUILDS FROM BOT
func (m *Model) GetBotGuilds(guildID string) ([]Guild, error) {
	bot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrCreateDiscordGoSession, err)
	}

	if err := bot.Open(); err != nil {
		return nil, errors.ErrOpeningDiscordConnection
    }
	defer bot.Close()

	guilds, err := bot.UserGuilds(100, "", "", false)
	if err != nil {
		return nil, err
	}

	var userGuilds []Guild
	for _, g := range guilds {
		userGuilds = append(userGuilds, Guild{Id: g.ID, Name: g.Name})
	}

	return userGuilds, nil
}

// * GET GUILDS FROM USER
func (m *Model) GetUserGuilds(discordToken string) ([]Guild, error) {
	userSession, err := discordgo.New("Bearer " + discordToken)
	if err != nil {
		return nil, err
	}

	guilds, err := userSession.UserGuilds(100, "", "", false)
	if err != nil {
		return nil, err
	}

	var userGuilds []Guild
	for _, g := range guilds {
		userGuilds = append(userGuilds, Guild{Id: g.ID, Name: g.Name})
	}
	return userGuilds, nil
}

// * FIND COMMON GUILDS USER-BOT
func (m *Model) FindCommonGuilds(botGuilds, userGuilds []Guild) []Guild {
	botGuildMap := make(map[string]Guild)
	for _, guild := range botGuilds {
		botGuildMap[guild.Id] = guild
	}

	var commonGuilds []Guild
	for _, guild := range userGuilds {
		if _, exists := botGuildMap[guild.Id]; exists {
			commonGuilds = append(commonGuilds, guild)
		}
	}

	return commonGuilds
}

func (m *Model) AddSession(data *DiscordSessionModel) error {
	ctx := context.TODO()
	newSync := DiscordSessionModel{
		UserId:  data.UserId,
		DiscordId: data.DiscordId,
		GuildId: data.GuildId,
		ChannelId: data.ChannelId,
		ActionId: data.ActionId,
	}

	_, err := m.Collection.InsertOne(ctx, newSync)
	if err != nil {
		return errors.ErrAddDiscordSession
	}

	log.Printf("Discord session created for user %s...\n", data.UserId)

	return nil
}

// func (m *Model) UpdateSession(userId string, session *UpdateDiscordSessionModel) error {
// 	ctx := context.TODO()
// 	filter := bson.M{"user_id": userId}
// 	update := bson.M{"$set": bson.M{"running": session.Running, "stop": session.Stop}}

// 	_, err := m.Collection.UpdateOne(
// 		ctx,
// 		filter,
// 		update,
// 	)
// 	if err != nil {
// 		return errors.ErrUpdateDiscordSession
// 	}

// 	return nil
// }

func  (m *Model) GetSessionByUserId(userId string) (*DiscordSessionModel, error) {
	var discordSession DiscordSessionModel
	err := m.Collection.FindOne(context.TODO(), bson.M{"user_id": userId}).Decode(&discordSession)
	if err != nil {
		return nil, errors.ErrDiscordUserSessionNotFound
	}

	return &discordSession, nil
}


func  (m *Model) GetAllDiscordSessions() ([]DiscordSessionModel, error) {
	ctx := context.TODO()
	cursor, err := m.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.ErrDiscordUserSessionNotFound
	}

	var discordSessions []DiscordSessionModel
	if err = cursor.All(ctx, &discordSessions); err != nil {
		return nil, errors.ErrDiscordUserSessionNotFound
	}

	return discordSessions, nil
}

// func (m *Model) GetSessionByDiscordId(discordId string) (*DiscordSessionModel, error) {
// 	var discordSession DiscordSessionModel
// 	err := m.Collection.FindOne(context.TODO(), bson.M{"discord_id": discordId}).Decode(&discordSession)
// 	if err != nil {
// 		return nil, errors.ErrDiscordUserSessionNotFound
// 	}

// 	return &discordSession, nil
// }

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