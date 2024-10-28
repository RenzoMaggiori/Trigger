package trigger

import (
	// "context"
	"context"
	"log"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/discord"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/errors"

	"trigger.com/trigger/pkg/decode"
)

func (h *Handler) WatchDiscord(w http.ResponseWriter, r *http.Request) {
	log.Printf("Watching discord")
	accessToken := r.Header.Get("Authorization")
	actionNode, err := decode.Json[workspace.ActionNodeModel](r.Body)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	err = h.Service.Watch(context.WithValue(context.TODO(), discord.AccessTokenCtxKey, accessToken), actionNode)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}

func (h *Handler) WebhookDiscord(w http.ResponseWriter, r *http.Request) {
	log.Printf("Webhook discord")
	// event, err := decode.Json[discord.Event](r.Body)

	// if err != nil {
	// 	customerror.Send(w, err, errors.ErrCodes)
	// }

	// log.Printf("Webhook triggered, received body=%+v\n", event)

	err := h.Service.Webhook(context.TODO())

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) StopDiscord(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("Authorization")

	err := h.Service.Stop(context.WithValue(context.TODO(), discord.AccessTokenCtxKey, accessToken))

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}

func (h *Handler) GetGuilds(w http.ResponseWriter, r *http.Request) {
	err := h.Service.Guilds()

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}

func (h *Handler) GetGuildChannels(w http.ResponseWriter, r *http.Request) {
	guildID := r.PathValue("guild_id")

	if guildID == "" {
		customerror.Send(w, errors.ErrUserTypeNone, errors.ErrCodes)
		return
	}

	err := h.Service.GuildChannels(guildID)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}