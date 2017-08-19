package main

import (
	"context"
	"fmt"
	// "net/http"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/google/go-github/github"
	// "github.com/nlopes/slack"
	"golang.org/x/oauth2"
)

type gpsiConfig struct {
	GitHub githubConfig
	Slack  slackConfig
}

type githubConfig struct {
	Token string
	Owner string
	Repos []string
    Projects []string
}

type slackConfig struct {
	Token    string
	User     string
	Team     string
	Channels []string
}

func main() {
	config := gpsiConfig{}
	meta, err := toml.DecodeFile("gpsi.toml", &config)
	if &meta == nil && err != nil {
		fmt.Println(err)
	}

	githubToken := config.GitHub.Token
	githubOwner := config.GitHub.Owner
	githubRepos := config.GitHub.Repos
    // githubProjects := config.GitHub.Projects

	// slackToken := config.Slack.Token
	// slackUser := config.Slack.User
	// slackTeam := config.Slack.Team
	// slackChannels := config.Slack.Channels

	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: githubToken,
	})
	authClient := oauth2.NewClient(ctx, tokenSource)
	githubClient := github.NewClient(authClient)

    fmt.Printf("STATUS - establishing client: %v", githubClient.UserAgent)

	createdAt := time.Now()
	hookName := "gpsi"
	hookURL := "http://localhost:7000"
	// https://developer.github.com/v3/activity/events/types/#projectcardevent
	events := []string{""}
	hookID := 42

	githubHook := github.Hook{
		CreatedAt: &createdAt,
		Name:      &hookName,
		URL:       &hookURL,
		Events:    events,
		ID:        &hookID,
	}

	for i := 0; i < len(githubRepos); i++ {
		hook, resp, err := githubClient.Repositories.CreateHook(ctx, githubOwner, githubRepos[i], &githubHook)
		if err != nil {
            fmt.Printf("ERROR - %v hook placement: %v, response status: %v", *hook.Name, err, resp.Status)
		}
	}

	// slackClient := slack.New(slackToken)

	// hookHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// })

    // Outline:
    // [ ] if no command line arguments exists
    // [X] - read from gpsi.toml file
    // [ ] else
    // [ ] - read from command line arguments
    // [X] create slack github client
    // [X] create slack client
    // [ ] place hooks on target repositories
    // [ ] if hook exists
    // [ ] - report it exists and move to next hook
    // [ ] create listening server
    // [ ] - start listenandserve
    // [ ] create hook handler
    // [ ] - listen for specific event lists
    // [ ] - create new slack message
    // [ ] - post to target channel(s)

    // Resources:
    // https://github.com/google/go-github/blob/master/github/projects.go
    // https://github.com/google/go-github/blob/master/github/repos_projects.go
}
