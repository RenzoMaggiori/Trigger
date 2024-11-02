package workspace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func StartActionRequest(accessToken string, actionCompletedNode ActionNodeModel, action action.ActionModel) (int, error) {
	body, err := json.Marshal(actionCompletedNode)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	actionEnv := fmt.Sprintf("%s_SERVICE_BASE_URL", strings.ToUpper(action.Provider))
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/%s/%s/%s", os.Getenv(actionEnv), action.Provider, action.Type, action.Action),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)

	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return res.StatusCode, errors.ErrSettingAction
	}
	return res.StatusCode, nil
}

func StopActionRequest(accessToken string, workspace *WorkspaceModel, actionNode ActionNodeModel, action action.ActionModel) (int, error) {
	body, err := json.Marshal(actionNode)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/%s/trigger/stop", os.Getenv("ACTION_SERVICE_BASE_URL"), action.Provider),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return res.StatusCode, errors.ErrSettingAction
	}
	return res.StatusCode, nil
}

func ActionCompletedRequest(accessToken string, update ActionCompletedModel) (int, error) {
	body, err := json.Marshal(update)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPatch,
			fmt.Sprintf("%s/api/workspace/action_completed", os.Getenv("ACTION_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
				"Content-Type":  "application/json",
			},
		),
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return res.StatusCode, errors.ErrCompletingAction
	}
	return res.StatusCode, nil
}

func GetByUserId(accessToken string, userId string) ([]WorkspaceModel, int, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/workspace/user_id/%s", os.Getenv("ACTION_SERVICE_BASE_URL"), userId),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
				"Content-Type":  "application/json",
			},
		),
	)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrCompletingAction
	}

	workspaces, err := decode.Json[[]WorkspaceModel](res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	return workspaces, res.StatusCode, nil
}

func GetByActionIdRequest(accessToken string, actionId string) ([]WorkspaceModel, int, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/workspace/action_id/%s", os.Getenv("ACTION_SERVICE_BASE_URL"), actionId),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
				"Content-Type":  "application/json",
			},
		),
	)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrCompletingAction
	}

	workspaces, err := decode.Json[[]WorkspaceModel](res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}
	return workspaces, res.StatusCode, nil
}

func WatchCompletedRequest(accessToken string, watchCompleted WatchCompletedModel) (int, error) {
	body, err := json.Marshal(watchCompleted)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPatch,
			fmt.Sprintf("%s/api/workspace/watch_completed", os.Getenv("ACTION_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
				"Content-Type":  "application/json",
			},
		),
	)

	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return res.StatusCode, errors.ErrCompletingAction
	}

	return res.StatusCode, nil
}
