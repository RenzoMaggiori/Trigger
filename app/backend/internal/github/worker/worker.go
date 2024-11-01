package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
		if err := changeInPushes(ctx); err != nil {
			log.Println(err)
		}
		log.Println("job ended")
	})
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func changeInPushes(ctx context.Context) error {
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
	// fmt.Printf("Got %d workspaces with github actions in it", len(githubWorkspaces))

	for _, workspace := range githubWorkspaces {
		log.Printf("Github checking for changes in commit in workspace: %s", workspace.Id.Hex())
		err = userChangeInPushes(ctx, workspace, *githubAction)
		if err != nil {
			log.Printf("Error while checking changes in commits in workspace[%s]: %s", workspace.Id.Hex(), err)
		}
	}
	return nil
}

func userChangeInPushes(ctx context.Context, workspace workspace.WorkspaceModel, githubAction action.ActionModel) error {

	session, _, err := session.GetSessionByUserIdRequest(os.Getenv("ADMIN_TOKEN"), workspace.UserId.Hex())

	if err != nil {
		return err
	}

	// log.Printf("Github commit watch got session: %+v", session)

	sync, _, err := sync.GetSyncAccessTokenRequest(session[0].AccessToken, workspace.UserId.Hex(), "github")

	if err != nil {
		return err
	}

	// log.Printf("Github commit watch got sync: %+v", sync)

	client := github.NewClient(nil).WithAuthToken(sync.AccessToken)

	githubUser, res, err := client.Users.Get(ctx, "")

	if err != nil {
		return err
	}

	if res.StatusCode > 200 {
		return errors.ErrGithubUserInfo
	}

	log.Printf("Github commit watch got github user: %+v", githubUser)

	//TODO: Create a go routine for the loop :)
	for _, node := range workspace.Nodes {
		if node.ActionId == githubAction.Id && node.Status == "active" {
			log.Printf("Checking node[%s]", node.NodeId)
			since, err := time.Parse("2006-01-02 15:04:05", node.Input["since"])
			// TODO: not returning, just adding error into a channel
			if err != nil {
				return err
			}
			commit, res, err := client.Repositories.ListCommits(ctx, *githubUser.Login, node.Input["repo"], &github.CommitsListOptions{
				Since: since,
			})

			if err != nil {
				return err
			}

			if res.StatusCode >= 400 {
				return errors.ErrGithubUserInfo
			}

			log.Printf("Got %d commits", len(commit))

			if len(commit) > 0 {
				sendCommitWebhook(sync.AccessToken, commit[len(commit)-1])
			}
		}
	}
	return nil
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
