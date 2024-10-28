package worker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/spotify"
	"trigger.com/trigger/internal/spotify/trigger"
	userSync "trigger.com/trigger/internal/sync"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/mongodb"
)

var (
	errCollectionNotFound error = errors.New("could not find spotify collection")
	errSessionNotFound    error = errors.New("could not find user session")
	errSyncModelNull      error = errors.New("the sync models type is null")
	errSpotifyAction      error = errors.New("spotify action not found")
	errSpotifyBadStatus   error = errors.New("bad status code from spotify")
	errWebhookBadStatus   error = errors.New("webhook returned a bad status")
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

	var wg sync.WaitGroup
	for _, w := range workspaces {
		wg.Add(1)
		go func(userId string) {
			defer wg.Done()
			err := userChangeInFollowers(ctx, collection, userId)
			if err != nil {
				log.Printf("Error processing user %s: %v", userId, err)
			}
		}(w.UserId.String())
	}
	wg.Wait()
	return nil
}

func userChangeInFollowers(ctx context.Context, collection *mongo.Collection, userId string) error {
	accessToken, err := getUserAccessToken(userId)
	if err != nil {
		return err
	}

	spotifyUser, err := getSpotifyUser(accessToken)
	if err != nil {
		return err
	}

	var userHistory SpotifyFollowerHistory
	filter := bson.M{"user_id": userId}
	if err = collection.FindOne(ctx, filter).Decode(&userHistory); err != nil {
		if err == mongo.ErrNoDocuments {
			spotifyHistory := SpotifyFollowerHistory{
				UserId: userId,
				Total:  spotifyUser.Followers.Total,
			}
			_, err := collection.InsertOne(ctx, spotifyHistory)
			if err != nil {
				return err
			}
			userHistory.Total = spotifyUser.Followers.Total
		} else {
			return err
		}
	}

	if spotifyUser.Followers.Total != userHistory.Total {
		err := fetchSpotifyWebhook(trigger.FollowerChange{
			Followers: spotifyUser.Followers.Total,
			Increased: spotifyUser.Followers.Total > userHistory.Total,
		})
		if err != nil {
			return err
		}

		_, err = collection.UpdateOne(ctx, filter, bson.M{
			"$set": bson.M{"total": spotifyUser.Followers.Total},
		},
		)
		if err != nil {
			return err
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
	if spotifyFollowerAction == "" {
		return nil, errSpotifyAction
	}

	workspaces, _, err := workspace.GetByActionIdRequest(
		os.Getenv("ADMIN_TOKEN"),
		spotifyFollowerAction,
	)
	return workspaces, err
}

func getUserAccessToken(userId string) (string, error) {
	session, _, err := session.GetSessionByUserIdRequest(os.Getenv("ADMIN_TOKEN"), userId)
	if err != nil {
		return "", err
	}
	if session == nil || len(session) == 0 {
		return "", errSessionNotFound
	}

	user, _, err := userSync.GetSyncAccessTokenRequest(session[0].AccessToken, userId, "spotify")
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errSyncModelNull
	}
	return user.AccessToken, nil
}

func getSpotifyUser(accessToken string) (*SpotifyUser, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/me", spotify.BaseUrl),
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
	if res.StatusCode != http.StatusOK {
		return nil, errSpotifyBadStatus
	}

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

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/spotify/trigger/webhook", os.Getenv("SPOTIFY_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			nil,
		),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errWebhookBadStatus
	}
	return nil
}
