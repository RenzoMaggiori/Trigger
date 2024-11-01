package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	goSync "sync"
	"time"

	"github.com/google/go-github/v66/github"
	"github.com/robfig/cron"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/sync"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func New(ctx context.Context) *cron.Cron {
	c := cron.New()
	err := c.AddFunc("0 */1 * * * *", func() {
		log.Println("job running...")
		if err := changeInCommits(ctx); err != nil {
			log.Println(err)
		}
		log.Println("job ended")
	})
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func changeInCommits(ctx context.Context) error {
	githubAction, err := getGithubAction()

	if err != nil {
		return err
	}

	githubWorkspaces, _, err := workspace.GetByActionIdRequest(
		os.Getenv("ADMIN_TOKEN"),
		githubAction.Id.Hex())

	if err != nil {
		return err
	}

	for _, workspace := range githubWorkspaces {
		err = userChangeInCommits(ctx, workspace, *githubAction)
		if err != nil {
			return err
		}
	}
	return nil
}

func userChangeInCommits(ctx context.Context, userWorkspace workspace.WorkspaceModel, githubAction action.ActionModel) error {
	// Fetch session and sync access token
	accessToken, syncToken, err := getSessionAndSyncToken(userWorkspace)
	if err != nil {
		return err
	}
	// Initialize GitHub client
	client := github.NewClient(nil).WithAuthToken(syncToken)
	githubUser, err := fetchGithubUser(ctx, client)
	if err != nil {
		return err
	}

	var wg goSync.WaitGroup
	errChan := make(chan error, len(userWorkspace.Nodes))

	for _, node := range userWorkspace.Nodes {
		if node.ActionId == githubAction.Id && node.Status == "active" {
			wg.Add(1)
			go nodeChangeInCommits(ctx, &wg, errChan, client, githubUser, node, accessToken)
		}
	}

	// Wait for all goroutines to finish and check for errors
	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

func getSessionAndSyncToken(userWorkspace workspace.WorkspaceModel) (string, string, error) {
	session, _, err := session.GetSessionByUserIdRequest(os.Getenv("ADMIN_TOKEN"), userWorkspace.UserId.Hex())
	if err != nil {
		return "", "", err
	}
	sync, _, err := sync.GetSyncAccessTokenRequest(session[0].AccessToken, userWorkspace.UserId.Hex(), "github")
	if err != nil {
		return "", "", err
	}
	return session[0].AccessToken, sync.AccessToken, nil
}

func fetchGithubUser(ctx context.Context, client *github.Client) (*github.User, error) {
	githubUser, res, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 200 {
		return nil, errors.ErrGithubUserInfo
	}
	return githubUser, nil
}

func nodeChangeInCommits(
	ctx context.Context,
	wg *goSync.WaitGroup,
	errChan chan<- error,
	client *github.Client,
	githubUser *github.User,
	node workspace.ActionNodeModel,
	accessToken string,
) {
	defer wg.Done()

	since, err := time.Parse("2006-01-02 15:04:05", node.Input["since"])
	if err != nil {
		errChan <- err
		return
	}

	commit, res, err := client.Repositories.ListCommits(ctx, *githubUser.Login, node.Input["repo"], &github.CommitsListOptions{Since: since})
	if err != nil {
		errChan <- err
		return
	}

	if res.StatusCode >= 400 {
		errChan <- errors.ErrGithubCommitInfo
	}

	log.Printf("Got %d commits", len(commit))
	if len(commit) > 0 {
		sendCommitWebhook(accessToken, commit[len(commit)-1])
	}
}

func getGithubAction() (*action.ActionModel, error) {
	actions, _, err := action.GetByProviderRequest(
		os.Getenv("ADMIN_TOKEN"),
		"github",
	)
	if err != nil {
		return nil, err
	}

	for _, a := range actions {
		if a.Type != "trigger" {
			continue
		}
		if a.Action != "watch_commit" {
			continue
		}
		return &a, nil
	}
	return nil, errors.ErrActionNotFound
}

func sendCommitWebhook(accessToken string, commit *github.RepositoryCommit) error {

	body, err := json.Marshal(commit)

	if err != nil {
		return err
	}

	res, err := fetch.Fetch(http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/github/trigger/webhook", os.Getenv("GITHUB_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)

	if err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		return errors.ErrGithubSendingWebhook
	}
	return nil
}
