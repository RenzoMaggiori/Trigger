package worker

import (
	"sync"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/internal/discord/trigger"
	"trigger.com/trigger/internal/session"
	userSync "trigger.com/trigger/internal/sync"
)

const (
	authURL      string = "https://discord.com/api/oauth2/authorize"
	tokenURL     string = "https://discord.com/api/v10/oauth2/token"
	userEndpoint string = "https://discord.com/api/v10/users/@me"
	baseURL      string = "https://discord.com/api/v10"

	workerBaseURL string = "http://localhost:8010/api/discord/worker"
)

type Service interface {
	GetGuildChannels(guildID string) ([]Channel, error)
	GetBotGuilds(guildID string) ([]Guild, error)
	GetUserGuilds(discordToken string) ([]Guild, error)
	FindCommonGuilds(botGuilds, userGuilds []Guild) []Guild

	AddSession(data *DiscordSessionModel) error
	GetSessionByUserId(userId string) (*DiscordSessionModel, error)
	GetAllDiscordSessions() ([]DiscordSessionModel, error)
	DeleteSession(userId string) error

	FetchDiscordWebhook(accessToken string, data trigger.ActionBody) error
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
	bot    *discordgo.Session
	mutex      sync.Mutex
}

type Guild struct {
	Id   string
	Name string
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserTokens struct {
	session session.SessionModel
	sync    userSync.SyncModel
}

type DiscordSessionModel struct {
	UserId  string `json:"user_id" bson:"user_id"`
	DiscordId string `json:"discord_id" bson:"discord_id"`
	GuildId string `json:"guild_id" bson:"guild_id"`
	ChannelId string `json:"channel_id" bson:"channel_id"`
	ActionId string `json:"action_id" bson:"action_id"`
	Token string `json:"token" bson:"token"`
}

// type WatchCompletedModel struct {
// 	UserId   primitive.ObjectID `json:"user_id"`
// 	ActionId primitive.ObjectID `json:"action_id"`
// 	Status   string             `json:"status"`
// 	Output   map[string]string  `json:"output"`
// }

type DiscordMe struct {
	DiscordId string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type DiscordStruct struct {
	Id                   string  `json:"id"`
	Username             string  `json:"username"`
	Avatar               *string `json:"avatar"`
	Discriminator        string  `json:"discriminator"`
	PublicFlags          int     `json:"public_flags"`
	Flags                int     `json:"flags"`
	Banner               *string `json:"banner"`
	AccentColor          *string `json:"accent_color"`
	GlobalName           string  `json:"global_name"`
	AvatarDecorationData *string `json:"avatar_decoration_data"`
	BannerColor          *string `json:"banner_color"`
	Clan                 *string `json:"clan"`
	MFAEnabled           bool    `json:"mfa_enabled"`
	Locale               string  `json:"locale"`
	PremiumType          int     `json:"premium_type"`
	Email                string  `json:"email"`
	Verified             bool    `json:"verified"`
}
