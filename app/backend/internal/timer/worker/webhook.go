package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/timer/trigger"
	"trigger.com/trigger/pkg/fetch"
)

func fetchWebhook(w workspace.WorkspaceModel, triggerName string, triggerType trigger.Type) error {
	session, _, err := session.GetSessionByUserIdRequest(os.Getenv("ADMIN_TOKEN"), w.UserId.Hex())
	if err != nil {
		return err
	}
	if len(session) == 0 {
		return errSessionNotFound
	}

	triggerAction := trigger.ActionBody{
		Type:     triggerType,
		Name:     triggerName,
		DateTime: time.Now(),
	}
	body, err := json.Marshal(triggerAction)
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/timer/trigger/webhook", os.Getenv("TIMER_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", session[0].AccessToken),
				"Content-Type":  "application/json",
			},
		),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return fmt.Errorf("%w: %s", errWebhookBadStatus, res.Status)
	}
	return nil
}
