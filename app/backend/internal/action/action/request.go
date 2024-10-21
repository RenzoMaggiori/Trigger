package action

import (
	"fmt"
	"net/http"
	"os"

	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func GetByIdRequest(accessToken string, actionId string) (*ActionModel, int, error) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/action/id/%s", os.Getenv("ACTION_SERVICE_BASE_URL"), actionId),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))

	if err != nil {
		return nil, res.StatusCode, errors.ErrFetchingActions
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrFetchingActions
	}
	action, err := decode.Json[ActionModel](res.Body)

	if err != nil {
		return nil, res.StatusCode, err
	}

	return &action, res.StatusCode, nil
}

func GetByProviderRequest(accessToken string, provider string) ([]ActionModel, int, error) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/action/provider/%s", os.Getenv("ACTION_SERVICE_BASE_URL"), provider),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))

	if err != nil {
		return nil, res.StatusCode, errors.ErrFetchingActions
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrFetchingActions
	}
	action, err := decode.Json[[]ActionModel](res.Body)

	if err != nil {
		return nil, res.StatusCode, err
	}

	return action, res.StatusCode, nil
}

func GetByActionNameRequest(accessToken string, actionName string) (*ActionModel, int, error) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/action/action/%s", os.Getenv("ACTION_SERVICE_BASE_URL"), actionName),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))

	if err != nil {
		return nil, res.StatusCode, errors.ErrFetchingActions
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrFetchingActions
	}
	action, err := decode.Json[ActionModel](res.Body)

	if err != nil {
		return nil, res.StatusCode, err
	}

	return &action, res.StatusCode, nil
}
