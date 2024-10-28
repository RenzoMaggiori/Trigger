package trigger

import "trigger.com/trigger/pkg/action"

type SpotifyWorkspaceCtx string

const SpotifyEventCtxKey SpotifyWorkspaceCtx = SpotifyWorkspaceCtx("SpotifyEventCtxKey")

type Service interface {
	action.Trigger
}

type Handler struct {
	Service
}

type Model struct {
}

type ActionBody struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type FollowerChange struct {
	Followers int  `json:"followers"`
	Increased bool `json:"increased"`
}
