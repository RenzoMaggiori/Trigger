package trigger

import (
	"sync"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/action"
)

type Service interface {
	action.Trigger
	Guilds() error
	GuildChannels(guildID string) error
}

type Handler struct {
	Service
}

// type Model struct {
// 	Collection *mongo.Collection
// }

type Model struct {
	Collection *mongo.Collection
    discord *discordgo.Session
    stop    chan struct{}
    running bool
    mutex   sync.Mutex
}

type StopModel struct {
	Owner  string `json:"owner"`
	Repo   string `json:"repo"`
	HookId string `json:"hookId"`
}
