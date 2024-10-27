package worker

import (
	"errors"
	"log"

	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/pkg/mongodb"
)

var (
	errCollectionNotFound = errors.New("could not find spotify collection")
)

func New(ctx context.Context) *cron.Cron {
	c := cron.New()
	err := c.AddFunc("0 0 * * *", func() {
		log.Println("job running...")
		if err := changeInFollowers(ctx); err != nil {
			log.Fatal(err)
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
	// TODO: Fetch users subscribed to Spotify (add query logic here)
	// Pseudo-code:
	// users, err := spotifyCollection.Find(ctx, bson.M{})
	// if err != nil { return fmt.Errorf("failed to fetch users: %w", err) }

	// Iterate over users
	// for each user:
	// 	// Get spotify user data using Spotify API client
	// 	newFollowerCount, err := FetchSpotifyFollowerCount(user.SpotifyID)
	// 	if err != nil {
	// 		log.Printf("Failed to fetch followers for user %s: %v", user.SpotifyID, err)
	// 		continue
	// 	}

	// 	// Compare with stored follower count
	// 	if newFollowerCount != user.StoredFollowerCount {
	// 		log.Printf("Follower count changed for user %s: %d -> %d", user.SpotifyID, user.StoredFollowerCount, newFollowerCount)
	// 		// Trigger action and update follower count in the database
	// 		user.StoredFollowerCount = newFollowerCount
	// 		_, err := spotifyCollection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": bson.M{"storedFollowerCount": newFollowerCount}})
	// 		if err != nil {
	// 			log.Printf("Failed to update follower count for user %s: %v", user.SpotifyID, err)
	// 		}
	// 	} else {
	// 		log.Printf("No change in follower count for user %s", user.SpotifyID)
	// 	}

	return nil
}

func getSpotifyWorkspaces() ([]workspace.WorkspaceModel, error) {
	return nil, nil
}
