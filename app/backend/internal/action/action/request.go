package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func GetActionByIdRequest(accessToken string, actionId string) (*action.ActionModel, string, error) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/action/id/%s", os.Getenv("ACTION_SERVICE_BASE_URL"),actionId),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))

	if err != nil {
		return nil, errors.errFetchingActions
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.errFetchingActions
	}
	action, err := decode.Json[action.ActionModel](res.Body)

	if err != nil {
		return nil, errors.errActionTypeNone
	}
	actionEnv := fmt.Sprintf("%s_SERVICE_BASE_URL", strings.ToUpper(action.Provider))

	return &action, actionEnv, nil
}

func ActionRequest(accessToken string, actionEnv string, action action.ActionModel, workspace *workspace.WorkspaceModel, node action.ActionNodeModel) error {
	body, err := json.Marshal(node)
	if err != nil {
		return err
	}
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
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.errAction
	}

	workspace.Nodes[0].Status = "active"

	return nil
}