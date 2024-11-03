package worker

import (
	"sync"

	"github.com/bwmarrin/discordgo"
	"trigger.com/trigger/internal/session"
	userSync "trigger.com/trigger/internal/sync"
)

type Service interface {}

type Handler struct {
	Service
}
type Model struct {
	bot    *discordgo.Session
	mutex      sync.Mutex
}

type UserTokens struct {
	session session.SessionModel
	sync    userSync.SyncModel
}