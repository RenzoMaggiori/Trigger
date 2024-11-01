package trigger

import (
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/action"
)

type WorkspaceCtx string
const DiscordEventCtxKey WorkspaceCtx = WorkspaceCtx("DiscordEventCtxKey")
const AccessTokenCtxKey WorkspaceCtx = WorkspaceCtx("AuthorizationCtxKey")
const WorkspaceCtxKey WorkspaceCtx = WorkspaceCtx("WorkspaceCtxKey")

const (
	authURL      string = "https://discord.com/api/oauth2/authorize"
	tokenURL     string = "https://discord.com/api/v10/oauth2/token"
	userEndpoint string = "https://discord.com/api/v10/users/@me"
	baseURL      string = "https://discord.com/api/v10"

	// workerBaseURL string = "http://localhost:8010/api/discord/worker"
)


type Service interface {
	action.Trigger
	GetAllSessions() ([]DiscordSessionModel, error)
}

type Handler struct {
	Service
}

type MsgInfo struct {
	Content string `json:"content"`
	AuthoId string `json:"author_id"`
	AuthoUsername string `json:"author_username"`
}

type Model struct {
	Collection *mongo.Collection
}

type ActionBody struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type DiscordSessionModel struct {
	UserId  string `json:"user_id" bson:"user_id"`
	ChannelId string `json:"channel_id" bson:"channel_id"`
	ActionId string `json:"action_id" bson:"action_id"`
	Token string `json:"token" bson:"token"`
	DiscordData *discordgo.User `json:"discord_data"`
}