package worker

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/timer/trigger"
)

func MinuteNotifier() TimerService {
	return MinuteService{
		ticker: time.NewTicker(time.Duration(1) * time.Minute),
	}
}

func (s MinuteService) Notify(ctx context.Context) error {
	adminToken := os.Getenv("ADMIN_TOKEN")
	actionName := "watch_minute"
	action, _, err := action.GetByActionNameRequest(adminToken, actionName)
	if err != nil {
		return err
	}

	workspaces, _, err := workspace.GetByActionIdRequest(adminToken, action.Id.Hex())
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, w := range workspaces {
		wg.Add(1)
		defer wg.Done()
		go func(ws workspace.WorkspaceModel) {
			if err := fetchWebhook(ws, actionName, trigger.MINUTE); err != nil {
				log.Println(err)
			}
		}(w)
	}
	wg.Wait()
	return nil
}

func (s MinuteService) Ticker() *time.Ticker {
	return s.ticker
}
