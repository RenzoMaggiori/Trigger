package trigger

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
)

func GetAllDiscordSessions() ([]DiscordSessionModel, error) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/discord/trigger/sessions", os.Getenv("DISCORD_SERVICE_BASE_URL")),
		nil,
		nil,
	))
	if err != nil {
		return nil, errors.New("could not fetch discord sessions")
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not fetch discord sessions")
	}

	session, err := decode.Json[[]DiscordSessionModel](res.Body)
	if err != nil {
		return nil, errors.New("could not decode discord sessions")
	}
	return session, nil
}