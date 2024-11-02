package worker

import "errors"

var (
	errWebhookBadStatus = errors.New("bad status from webhook")
	errSessionNotFound  = errors.New("session not found")
)
