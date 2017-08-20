package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/google/go-github/github"
	"github.com/nlopes/slack"
	"golang.org/x/oauth2"
)

type githubConfig struct {
	Token    string
	Owner    string
	Name     string
	Projects []string
}

type slackConfig struct {
	Token    string
	User     string
	Team     string
	Channels []string
}

type repoConfig struct {
	GitHub githubConfig
	Slack  slackConfig
}

type gpsiConfig struct {
	Repos []repoConfig
}

func main() {
	config := gpsiConfig{}
	meta, err := toml.DecodeFile("gpsi.toml", &config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("STATUS - metadata: %v\n", meta.Keys())

	reference := make(map[string]repoConfig)

	// NOTE: This is currently a workaround for the field in gpsiConfig being
	// stored as a slice rather than a map - this could be done if the TOML
	// file is stored as [repos.repo-name].
	for i := range config.Repos {
		reference[config.Repos[i].GitHub.Name] = config.Repos[i]
	}

	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: config.Repos[0].GitHub.Token,
	})
	githubClient := github.NewClient(oauth2.NewClient(ctx, tokenSource))
	fmt.Printf("STATUS - establishing client: %v\n", githubClient.UserAgent)

	createdAt := time.Now()
	githubHook := github.Hook{
		CreatedAt: &createdAt,
		Name:      github.String("gpsi"),
		URL:       github.String("http://localhost:7000/gpishook"), // TODO: Enable dynamically finding an open port.
		Events:    []string{"project_card"},
		ID:        github.Int(42), // TODO: This should likely be something dynamically assigned.
	}

	// NOTE: https://developer.github.com/v3/activity/events/types/#projectcardevent

	for i := range config.Repos {
		owner := config.Repos[i].GitHub.Owner
		repo := config.Repos[i].GitHub.Name
		hook, resp, err := githubClient.Repositories.CreateHook(ctx, owner, repo, &githubHook)
		if resp == nil && err != nil {
			fmt.Printf("ERROR - %v hook placement: %v, response status: %v\n", *hook.Name, err, resp.Status)
		}
	}

	slackClient := slack.New(config.Repos[0].Slack.Token)

    // channels, err := slackClient.GetChannels(true)



	hookHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// NOTE: I'm not sure this is how to reference that part of the JSON message.
		repo := r.FormValue("repository.name")
		action := r.FormValue("action")
		// url := r.FormValue("project_card.url")

		// channels := reference[repo].Slack.Channels

		params := slack.PostMessageParameters{}

		// for i := range channels {
		// 	slackClient.PostMessage(channels[i])
		// }

	})

	http.HandleFunc("/gpsihook", hookHandler)
	http.ListenAndServe(":7000", nil)

	// Outline:
	// [ ] if no command line arguments exists
	// [X] - read from gpsi.toml file
	// [ ] else
	// [ ] - read from command line arguments
	// [X] create slack github client
	// [X] create slack client
	// [X] place hooks on target repositories
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
