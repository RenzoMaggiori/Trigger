package trigger

import (
	"sync"

	"github.com/bwmarrin/discordgo"
	// "trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/pkg/action"
)

type Service interface {
	action.Trigger

	// Stop(ctx context.Context, userID string) error
	// Watch(ctx context.Context, userID string, actionNode workspace.ActionNodeModel) error
	// Webhook(ctx context.Context, userID string) error
	runDiscordSession(discord *discordgo.Session, stopChan chan struct{})
}

type Handler struct {
	Service
}

// type Model struct {
// 	Collection *mongo.Collection
// }

// type Model struct {
// 	Collection *mongo.Collection
//     discord *discordgo.Session
//     stop    chan struct{}
//     running bool
//     mutex   sync.Mutex
// }

type Model struct {
	// Collection *mongo.Collection
	discord    *discordgo.Session
	mutex      sync.Mutex
}

// type DiscordSessionModel struct {
// 	UserID  string `json:"user_id" bson:"user_id"`
// 	GuildId string `json:"guild_id" bson:"guild_id"`
// 	Token   string `json:"token" bson:"token"`
// 	Running bool   `json:"running" bson:"running"`
// 	Stop    bool   `json:"stop" bson:"stop"`
// }

type StopModel struct {
	Owner  string `json:"owner"`
	Repo   string `json:"repo"`
	HookId string `json:"hookId"`
}

type ActionBody struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
