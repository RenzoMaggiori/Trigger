package reaction

import (
	"trigger.com/trigger/pkg/action"
)

type Service interface {
	action.MultipleReactions
}

type Handler struct {
	Service
}

type Model struct {
}

type Guild struct {
	ID                       string   `json:"id"`
	Name                     string   `json:"name"`
	Icon                     string   `json:"icon"`
	Banner                   string   `json:"banner"`
	Owner                    bool     `json:"owner"`
	Permissions              string   `json:"permissions"`
	Features                 []string `json:"features"`
	ApproximateMemberCount   int      `json:"approximate_member_count"`
	ApproximatePresenceCount int      `json:"approximate_presence_count"`
}

type WebhookUser struct {
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	ID            string `json:"id"`
	Avatar        string `json:"avatar"`
	PublicFlags   int    `json:"public_flags"`
}
type Webhook struct {
	Name          string      `json:"name"`
	Type          int         `json:"type"`
	ChannelID     string      `json:"channel_id"`
	Token         string      `json:"token"`
	Avatar        *string     `json:"avatar"`
	GuildID       string      `json:"guild_id"`
	ID            string      `json:"id"`
	ApplicationID *string     `json:"application_id"`
	User          WebhookUser `json:"user"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type NewWebhook struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"` // [data:image/jpeg;base64,BASE64_ENCODED_JPEG_IMAGE_DATA] data:image/svg+xml;utf8
}

type NewMessage struct {
	Content string `json:"content,omitempty"`
	// Embeds          []*MessageEmbed         `json:"embeds"`
	TTS bool `json:"tts"`
	// Components      []MessageComponent      `json:"components"`
	// Files           []*File                 `json:"-"`
	// AllowedMentions *MessageAllowedMentions `json:"allowed_mentions,omitempty"`
	// Reference       *MessageReference       `json:"message_reference,omitempty"`
	StickerIDs []string `json:"sticker_ids"`
	// Flags           MessageFlags            `json:"flags,omitempty"`
	ChannelId string `json:"channel_id"`
}

type MessagegContent struct {
	Content string `json:"content,omitempty"`
	// Embeds          []*MessageEmbed         `json:"embeds"`
	TTS bool `json:"tts"`
	// Components      []MessageComponent      `json:"components"`
	// Files           []*File                 `json:"-"`
	// AllowedMentions *MessageAllowedMentions `json:"allowed_mentions,omitempty"`
	// Reference       *MessageReference       `json:"message_reference,omitempty"`
	// StickerIDs []string `json:"sticker_ids"`
	// Flags           MessageFlags            `json:"flags,omitempty"`
}
