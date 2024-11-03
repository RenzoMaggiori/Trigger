package trigger

import (
	"trigger.com/trigger/pkg/action"
)

type BitbucketWorkspaceCtx string

const userIdCtxKey BitbucketWorkspaceCtx = BitbucketWorkspaceCtx("userIdCtxKey")

const WebhookEventCtxKey BitbucketWorkspaceCtx = BitbucketWorkspaceCtx("WebhookEventCtxKey")

type WatchBody struct {
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Active      bool     `json:"active"`
	Secret      string   `json:"secret"`
	Events      []string `json:"events"`
}

type WebhookRequest struct {
	PullRequest *PullRequest `json:"pullrequest,omitempty"`
	Push        any          `json:"push,omitempty"`
}

type PullRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Service interface {
	action.Trigger
}

type Handler struct {
	Service
}

type Model struct {
}
