package worker

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/internal/action/action"
)

// WARNING: Action names must be unique

var (
	actions = [...]action.AddActionModel{
		{
			Provider: "gmail",
			Type:     "trigger",
			Action:   "watch",
			Input:    []string{},
			Output:   []string{},
		},
		{
			Provider: "gmail",
			Type:     "reaction",
			Action:   "send_email",
			Input:    []string{},
			Output:   []string{},
		},
		{
			Provider: "github",
			Type:     "trigger",
			Action:   "watch_push",
			Input:    []string{"owner", "repo"},
			Output:   []string{},
		},
		{
			Provider: "github",
			Type:     "reaction",
			Action:   "create_issue",
			Input:    []string{"owner", "repo", "title", "description"},
			Output:   []string{},
		},
		{
			Provider: "spotify",
			Type:     "trigger",
			Action:   "watch_followers",
			Input:    []string{},
			Output:   []string{"followers", "increased"},
		},
		{
			Provider: "spotify",
			Type:     "reaction",
			Action:   "play_music",
			Input:    []string{},
			Output:   []string{},
		},
		{
			Provider: "timer",
			Type:     "trigger",
			Action:   "watch_minute",
			Input:    []string{},
			Output:   []string{"datetime"},
		},
		{
			Provider: "timer",
			Type:     "trigger",
			Action:   "watch_hour",
			Input:    []string{},
			Output:   []string{"datetime"},
		},
		{
			Provider: "timer",
			Type:     "trigger",
			Action:   "watch_day",
			Input:    []string{},
			Output:   []string{"datetime"},
		},
		{
			Provider: "discord",
			Type:     "trigger",
			Action:   "watch_message",
			Input:    []string{},
			Output:   []string{"followers", "increased"},
		},
		{
			Provider: "discord",
			Type:     "reaction",
			Action:   "send_message",
			Input:    []string{},
			Output:   []string{},
		},
	}
)

func Run(collection *mongo.Collection) error {
	ctx := context.TODO()
	newActions := make([]interface{}, 0)
	for _, a := range actions {
		filter := bson.M{
			"provider": a.Provider,
			"type":     a.Type,
			"action":   a.Action,
		}
		err := collection.FindOne(ctx, filter)
		if err.Err() == nil {
			continue
		}

		newActions = append(newActions, a)
	}
	if len(newActions) == 0 {
		return nil
	}

	_, err := collection.InsertMany(ctx, newActions)
	return err
}
