package trigger

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/action"
)

type WorkspaceCtx string

const WorkspaceCtxKey WorkspaceCtx = WorkspaceCtx("WorkspaceCtxKey")

const WebhookVerificationCtxKey WorkspaceCtx = WorkspaceCtx("WebhookVerificationCtxKey")

const WebhookUserIdCtxKey WorkspaceCtx = WorkspaceCtx("WebhookUserIdCtxKey")

const VerificationMessageType = "webhook_callback_verification"

type Service interface {
	action.Trigger
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
}

type WebhookVerificationResponse struct {
	Challenge string `json:"challenge"`
}

type WebhookVerificationRequest struct {
	Challenge    string                   `json:"challenge"`
	Subscription VerificationSubscription `json:"subscription"`
}

type VerificationSubscription struct {
	ID        string                `json:"id"`
	Status    string                `json:"status"`
	Type      string                `json:"type"`
	Version   string                `json:"version"`
	Cost      int                   `json:"cost"`
	Condition VerificationCondition `json:"condition"`
	Transport VerificationTransport `json:"transport"`
	CreatedAt time.Time             `json:"created_at"`
}

type VerificationCondition struct {
	BroadcasterUserID string `json:"broadcaster_user_id"`
}

type VerificationTransport struct {
	Method   string `json:"method"`
	Callback string `json:"callback"`
}

type ChannelFollowCondition struct {
	BroadcasterUserID string `json:"broadcaster_user_id"`
	ModeratorUserID   string `json:"moderator_user_id"`
}

type ChannelFollowTransport struct {
	Method   string `json:"method"`
	Callback string `json:"callback"`
	Secret   string `json:"secret"`
}

type ChannelFollowSubscriptionBody struct {
	Type      string                 `json:"type"`
	Version   string                 `json:"version"`
	Condition ChannelFollowCondition `json:"condition"`
	Transport ChannelFollowTransport `json:"transport"`
}
