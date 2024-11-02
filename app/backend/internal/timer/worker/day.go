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

func DayNotifier() TimerService {
	return DayService{
		ticker: time.NewTicker(time.Duration(24) * time.Hour),
	}
}

func (s DayService) Notify(ctx context.Context) error {
	adminToken := os.Getenv("ADMIN_TOKEN")
	actionName := "watch_day"
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
			if err := fetchWebhook(ws, actionName, trigger.DAY); err != nil {
				log.Println(err)
			}
		}(w)
	}
	wg.Wait()
	return nil
}

func (s DayService) Ticker() *time.Ticker {
	return s.ticker
}
