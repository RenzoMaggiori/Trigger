package worker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	trigger "trigger.com/trigger/internal/spotify/trigger"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/mongodb"
)

const (
	spotifyBaseUrl string = "https://api.spotify.com/v1"
)

var (
	errCollectionNotFound = errors.New("could not find spotify collection")
)

func New(ctx context.Context) *cron.Cron {
	c := cron.New()
	err := c.AddFunc("*/5 * * * *", func() {
		log.Println("job running...")
		if err := changeInFollowers(ctx); err != nil {
			log.Println(err)
		}
		log.Println("job ended")
	})
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func changeInFollowers(ctx context.Context) error {
	collection, ok := ctx.Value(mongodb.CtxKey).(*mongo.Collection)
	if !ok || collection == nil {
		return errCollectionNotFound
	}

	workspaces, err := getSpotifyWorkspaces()
	if err != nil {
		return err
	}

	for _, w := range workspaces {
		accessToken, err := getUserAccessToken(w.UserId.String())
		if err != nil {
			return err
		}

		spotifyUser, err := getSpotifyUser(accessToken)
		if err != nil {
			return err
		}

		var userHistory SpotifyFollowerHistory
		filter := bson.M{"user_id": w.UserId.String()}
		if err = collection.FindOne(ctx, filter).Decode(&userHistory); err != nil {
			spotifyHistory := SpotifyFollowerHistory{
				UserId: w.UserId.String(),
				Total:  spotifyUser.Followers.Total,
			}
			if err == mongo.ErrNoDocuments {
				collection.InsertOne(ctx, spotifyHistory)
			} else {
				collection.UpdateOne(ctx, filter, spotifyHistory)
			}
			continue
		}

		if spotifyUser.Followers.Total != userHistory.Total {
			err := fetchSpotifyWebhook(trigger.FollowerChange{
				Followers: spotifyUser.Followers.Total,
				Increased: spotifyUser.Followers.Total > userHistory.Total,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getSpotifyWorkspaces() ([]workspace.WorkspaceModel, error) {
	actions, _, err := action.GetByProviderRequest(
		os.Getenv("ADMIN_TOKEN"),
		"spotify",
	)
	if err != nil {
		return nil, err
	}

	var spotifyFollowerAction string = ""
	for _, a := range actions {
		if a.Type != "trigger" {
			continue
		}
		if a.Action != "watch_followers" {
			continue
		}
		spotifyFollowerAction = a.Id.String()
	}

	workspaces, _, err := workspace.GetByActionIdRequest(
		os.Getenv("ADMIN_TOKEN"),
		spotifyFollowerAction,
	)
	return workspaces, err
}

func getUserAccessToken(userId string) (string, error) {
	// TODO: get the Access Token
	return "", nil
}

func getSpotifyUser(accessToken string) (*SpotifyUser, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/me", spotifyBaseUrl),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	spotifyUser, err := decode.Json[SpotifyUser](res.Body)
	if err != nil {
		return nil, err
	}

	return &spotifyUser, nil
}

func fetchSpotifyWebhook(followerChange trigger.FollowerChange) error {
	body, err := json.Marshal(followerChange)
	if err != nil {
		return err
	}

	_, err = fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/spotify/trigger/webhook", os.Getenv("SPOTIFY_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			nil,
		),
	)
	return err
}
