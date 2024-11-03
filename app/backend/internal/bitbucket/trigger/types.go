package trigger

import (
	"time"

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
	Issue Issue `json:"issue"`
}

type Issue struct {
	ID        int       `json:"id"`
	Component string    `json:"component"`
	Title     string    `json:"title"`
	Content   Content   `json:"content"`
	Priority  string    `json:"priority"`
	State     string    `json:"state"`
	Type      string    `json:"type"`
	Milestone NameField `json:"milestone"`
	Version   NameField `json:"version"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	Links     Links     `json:"links"`
}

type Content struct {
	Raw    string `json:"raw"`
	HTML   string `json:"html"`
	Markup string `json:"markup"`
}

type NameField struct {
	Name string `json:"name"`
}

type Links struct {
	Self Href `json:"self"`
	HTML Href `json:"html"`
}

type Href struct {
	Href string `json:"href"`
}

type Service interface {
	action.Trigger
}

type Handler struct {
	Service
}

type Model struct {
}
