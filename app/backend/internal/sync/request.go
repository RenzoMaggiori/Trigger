package sync

import (
	"fmt"
	"net/http"
	"os"

	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func GetSyncAccessTokenRequest(accessToken string, userId string, provider string) (*SyncModel, int, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/sync/%s/%s", os.Getenv("SYNC_SERVICE_BASE_URL"), userId, provider),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		))
	if err != nil {
		return nil, http.StatusInternalServerError, errors.ErrSyncAccessTokenNotFound
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrSyncAccessTokenNotFound
	}

	sync, err := decode.Json[SyncModel](res.Body)
	if err != nil {
		return nil, res.StatusCode, errors.ErrSessionTypeNone
	}
	return &sync, res.StatusCode, nil
}
