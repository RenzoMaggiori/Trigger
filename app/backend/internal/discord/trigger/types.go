package trigger

import (
	"sync"

	"github.com/bwmarrin/discordgo"
	// "trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/pkg/action"
)

type WorkspaceCtx string
const DiscordEventCtxKey WorkspaceCtx = WorkspaceCtx("DiscordEventCtxKey")
const AccessTokenCtxKey WorkspaceCtx = WorkspaceCtx("AuthorizationCtxKey")
const WorkspaceCtxKey WorkspaceCtx = WorkspaceCtx("WorkspaceCtxKey")


type Service interface {
	action.Trigger

	// Stop(ctx context.Context, userID string) error
	// Watch(ctx context.Context, userID string, actionNode workspace.ActionNodeModel) error
	// Webhook(ctx context.Context, userID string) error
}

type Handler struct {
	Service
}

type MsgInfo struct {
	Author string `json:"author"`
	Content string `json:"content"`
}

type Model struct {
}

type ActionBody struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type Event struct {
	GuildId   string `json:"guild_id"`
	ChannelId string `json:"channel_id"`
}