package trigger

import (
	"trigger.com/trigger/pkg/action"
)

type GithubWorkspaceCtx string

const GithubCommitCtxKey GithubWorkspaceCtx = GithubWorkspaceCtx("GithubCommitCtxKey")

type Service interface {
	action.Trigger
}

type Handler struct {
	Service
}

type Model struct {
}

type StopModel struct {
	Owner  string `json:"owner"`
	Repo   string `json:"repo"`
	HookId string `json:"hookId"`
}
