package worker

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

type Service interface {}

type Handler struct {
	Service
}

type Model struct {
	bot    *discordgo.Session
	mutex      sync.Mutex
}
