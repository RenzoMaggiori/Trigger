package reaction

import (
	"fmt"
	"net/http"
	"trigger.com/trigger/pkg/fetch"
)

func (h *Handler) Guilds(accessToken string) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s", baseURL, "users/@me/guilds"),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))

	if err != nil {
		return
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return
	}

}