package settings

import (
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/sync"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) GetById(id primitive.ObjectID) (*sync.SyncModel, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPatch,
			fmt.Sprintf("%s/api/sync/id/%s", os.Getenv("SYNC_SERVICE_BASE_URL"), id.Hex()),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v", err)
	}

	sync, err := decode.Json[sync.SyncModel](res.Body)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &sync, nil
}

func (m Model) GetByUserId(userId string) (*sync.SyncModel, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPatch,
			fmt.Sprintf("%s/api/sync/user/%s", os.Getenv("SYNC_SERVICE_BASE_URL"), userId),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v", err)
	}

	syncs, err := decode.Json[sync.SyncModel](res.Body)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &syncs, nil
}