import workspace

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func GetActionByIdRequest(accessToken string, actionId string) (*action.ActionModel, error) {
	for i, node := range workspace.Nodes {
		if node.Status == "pending" {
			res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
				http.MethodGet,
				fmt.Sprintf("%s/api/action/id/%s", os.Getenv("ACTION_SERVICE_BASE_URL"), node.ActionId.Hex()),
				nil,
				map[string]string{
					"Authorization": fmt.Sprintf("Bearer %s", accessToken),
				},
			))

			if err != nil {
				return nil, errFetchingActions
			}
			defer res.Body.Close()
			if res.StatusCode != http.StatusOK {
				return nil, errFetchingActions
			}
			action, err := decode.Json[action.ActionModel](res.Body)

			if err != nil {
				return nil, errActionTypeNone
			}
			actionEnv := fmt.Sprintf("%s_SERVICE_BASE_URL", strings.ToUpper(action.Provider))

			body, err := json.Marshal(node)

			if err != nil {
				return nil, err
			}
			// Call the reaction / trigger

			res, err = fetch.Fetch(
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
				return nil, err
			}
			defer res.Body.Close()
			if res.StatusCode != http.StatusOK {
				return nil, errAction
			}
			workspace.Nodes[i].Status = "active"
		}
	}
	return &workspace, nil
}

func GetActionByIdRequest(accessToken string, actionId string) (*action.ActionModel, error) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/action/id/%s", os.Getenv("ACTION_SERVICE_BASE_URL"),actionId),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))

	if err != nil {
		return nil, errFetchingActions
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errFetchingActions
	}
	action, err := decode.Json[action.ActionModel](res.Body)

	if err != nil {
		return nil, errActionTypeNone
	}
	actionEnv := fmt.Sprintf("%s_SERVICE_BASE_URL", strings.ToUpper(action.Provider))

	body, err := json.Marshal(node)

	if err != nil {
		return nil, err
	}

	return &action, nil

}

func ActionRequest(accessToken string, action action.ActionModel, workspace *WorkspaceModel) {
	res, err = fetch.Fetch(
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
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errAction
	}

	workspace.Nodes[i].Status = "active"
}