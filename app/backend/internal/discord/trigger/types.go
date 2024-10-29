package trigger

import (
	"context"
	"sync"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/internal/action/workspace"
	// "trigger.com/trigger/pkg/action"
)

type Service interface {
	// action.Trigger

	Stop(ctx context.Context, userID string) error
	Watch(ctx context.Context, userID string, actionNode workspace.ActionNodeModel) error
	Webhook(ctx context.Context, userID string) error
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
    Collection *mongo.Collection
    discord    *discordgo.Session
    mutex      sync.Mutex
}


type UserSession struct {
    UserID   string `bson:"user_id"`
    Token    string `bson:"token"`
    Running  bool   `bson:"running"`
    StopChan bool   `bson:"stop_chan"` // indicates if a stop channel is active
}


type StopModel struct {
	Owner  string `json:"owner"`
	Repo   string `json:"repo"`
	HookId string `json:"hookId"`
}
