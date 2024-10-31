package worker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	"trigger.com/trigger/pkg/auth/oaclient"
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
	err := c.AddFunc("0 */1 * * * *", func() {
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

	spotifyAction, err := getSpotifyAction()
	if err != nil {
		return err
	}

	workspaces, _, err := workspace.GetByActionIdRequest(
		os.Getenv("ADMIN_TOKEN"),
		spotifyAction.Id.Hex(),
	)
	if err != nil {
		return err
	}

	log.Printf("%v\n", workspaces)
	var wg sync.WaitGroup
	for _, w := range workspaces {
		wg.Add(1)
		go func(w workspace.WorkspaceModel, a action.ActionModel) {
			defer wg.Done()
			err := userChangeInFollowers(ctx, collection, w, a)
			if err != nil {
				log.Printf("Error processing user %s: %v", w.UserId.Hex(), err)
			}
		}(w, *spotifyAction)
	}
	wg.Wait()
	return nil
}

func userChangeInFollowers(ctx context.Context, collection *mongo.Collection, workspace workspace.WorkspaceModel, action action.ActionModel) error {
	userId := workspace.UserId.Hex()
	userTokens, err := getUserAccessToken(userId)
	if err != nil {
		return err
	}

	spotifyUser, err := getSpotifyUser(userTokens.sync)
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
		err := fetchSpotifyWebhook(userTokens.session.AccessToken, trigger.ActionBody{
			Type: action.Action,
			Data: trigger.FollowerChange{
				Followers: spotifyUser.Followers.Total,
				Increased: spotifyUser.Followers.Total > userHistory.Total,
			},
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

func getSpotifyAction() (*action.ActionModel, error) {
	actions, _, err := action.GetByProviderRequest(
		os.Getenv("ADMIN_TOKEN"),
		"spotify",
	)
	if err != nil {
		return nil, err
	}

	for _, a := range actions {
		if a.Type != "trigger" {
			continue
		}
		if a.Action != "watch_followers" {
			continue
		}
		return &a, nil
	}
	return nil, errSpotifyAction
}

func getUserAccessToken(userId string) (*UserTokens, error) {
	session, _, err := session.GetSessionByUserIdRequest(os.Getenv("ADMIN_TOKEN"), userId)
	if err != nil {
		return nil, err
	}
	if session == nil || len(session) == 0 {
		return nil, errSessionNotFound
	}

	userTokens := UserTokens{
		session: session[0],
	}
	syncModel, _, err := userSync.GetSyncAccessTokenRequest(userTokens.session.AccessToken, userId, "spotify")
	if err != nil {
		return nil, err
	}
	if syncModel == nil {
		return nil, errSyncModelNull
	}

	userTokens.sync = *syncModel
	return &userTokens, nil
}

func getSpotifyUser(syncModel userSync.SyncModel) (*SpotifyUser, error) {
	client, err := oaclient.New(context.TODO(), spotify.Config(), &syncModel)
	if err != nil {
		return nil, err
	}

	res, err := fetch.Fetch(
		client,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/me", spotify.BaseUrl),
			nil,
			nil,
		),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, errSpotifyBadStatus
		}
		return nil, fmt.Errorf("%w: %s", errSpotifyBadStatus, body)
	}

	spotifyUser, err := decode.Json[SpotifyUser](res.Body)
	if err != nil {
		return nil, err
	}

	return &spotifyUser, nil
}

func fetchSpotifyWebhook(accessToken string, data trigger.ActionBody) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/spotify/trigger/webhook", os.Getenv("SPOTIFY_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
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
