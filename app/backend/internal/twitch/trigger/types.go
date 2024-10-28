package trigger

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/action"
)

type WorkspaceCtx string

const AccessTokenCtxKey WorkspaceCtx = WorkspaceCtx("AuthorizationCtxKey")

const WorkspaceCtxKey WorkspaceCtx = WorkspaceCtx("WorkspaceCtxKey")

const WebhookVerificationCtxKey WorkspaceCtx = WorkspaceCtx("WebhookVerificationCtxKey")

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

type AppAccessTokenBody struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type AppAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
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

type UserResponse struct {
	Data []User `json:"data"`
}

type User struct {
	ID              string    `json:"id"`
	Login           string    `json:"login"`
	DisplayName     string    `json:"display_name"`
	Type            string    `json:"type"`
	BroadcasterType string    `json:"broadcaster_type"`
	Description     string    `json:"description"`
	ProfileImageURL string    `json:"profile_image_url"`
	OfflineImageURL string    `json:"offline_image_url"`
	ViewCount       int       `json:"view_count"`
	Email           string    `json:"email"`
	CreatedAt       time.Time `json:"created_at"`
}
