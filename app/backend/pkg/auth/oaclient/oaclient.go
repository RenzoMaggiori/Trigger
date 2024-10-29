package oaclient

import (
	"context"
	"errors"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"trigger.com/trigger/internal/sync"
)

var (
	errConfigNotFound = errors.New("could not find config")
	errSyncNotFound   = errors.New("could not find config")
)

func New(ctx context.Context, config *oauth2.Config, syncModel *sync.SyncModel) (*http.Client, error) {
	if config == nil {
		return nil, errConfigNotFound
	}
	if syncModel == nil {
		return nil, errSyncNotFound
	}

	refreshToken := ""
	if syncModel.RefreshToken != nil {
		refreshToken = *syncModel.RefreshToken
	}

	client := config.Client(ctx, &oauth2.Token{
		AccessToken:  syncModel.AccessToken,
		TokenType:    "Bearer",
		RefreshToken: refreshToken,
		Expiry:       syncModel.Expiry,
		ExpiresIn:    syncModel.Expiry.Unix() - time.Now().Unix(),
	})
	return client, nil
}
