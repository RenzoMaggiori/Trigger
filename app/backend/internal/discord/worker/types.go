package worker

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	authURL      string = "https://discord.com/api/oauth2/authorize"
	tokenURL     string = "https://discord.com/api/v10/oauth2/token"
	userEndpoint string = "https://discord.com/api/v10/users/@me"
	baseURL      string = "https://discord.com/api/v10"
)

type Service interface {
	Me(token string) (*Me, error)
	GuildChannels(guildID string) ([]Channel, error)
	AddSession(session *AddDiscordSession) error
	UpdateSession(userId string, session *UpdateDiscordSession) error
	GetSession(token string) (*DiscordSessionModel, error)
	DeleteSession(userId string) error
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
	// discord    *discordgo.Session
	// mutex      sync.Mutex
}

type DiscordSessionModel struct {
	UserId  string `json:"user_id" bson:"user_id"`
	DiscordId string `json:"discord_id" bson:"discord_id"`
	GuildId string `json:"guild_id" bson:"guild_id"`
	Token   string `json:"token" bson:"token"`
	Running bool   `json:"running" bson:"running"`
	Stop    bool   `json:"stop" bson:"stop"`
}

type AddDiscordSession struct {
	UserId  string `json:"user_id" bson:"user_id"`
	DiscordId string `json:"discord_id" bson:"discord_id"`
	GuildId string `json:"guild_id" bson:"guild_id"`
}

type UpdateDiscordSession struct {
	Running bool   `json:"running" bson:"running"`
	Stop    bool   `json:"stop" bson:"stop"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type Me struct {
	DiscordId string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type DiscordMe struct {
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
