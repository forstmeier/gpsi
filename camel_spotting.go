package camelspotting

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
	GitHub   string   // GitHub token
	User     string   // GitHub username
	Slack    string   // Slack token (or webhook URL)
	Team     string   // Slack team
	Channels []string // Slack channels
	Repo     []repoConfig
	Org      []orgConfig
}

type repoConfig struct {
	Owner    string
	Name     string
	Projects []string
	Channels []string // Slack specific override channels
}

type orgConfig struct {
	Org      string
	Name     string
	Projects []string
	Channels []string // Slack specific override channels
}

func main() {
	config := gpsiConfig{}
	meta, err := toml.DecodeFile("gpsi.toml", &config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("STATUS - metadata: %v\n", meta.Keys())

	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: config.GitHub,
	})
	githubClient := github.NewClient(oauth2.NewClient(ctx, tokenSource))
	fmt.Printf("STATUS - establishing client: %v\n", githubClient.UserAgent)

	createdAt := time.Now()
	githubHook := github.Hook{
		CreatedAt: &createdAt,
		Name:      github.String("gpsi"),
		URL:       github.String("http://localhost:7000/gpishook"), // TODO: Enable dynamically finding an open port.
		Events:    []string{"project_card"},
		// ID:        github.Int(42), // TODO: This should likely be something dynamically assigned.
	}
	// NOTE: https://developer.github.com/v3/activity/events/types/#projectcardevent

	repoReference := make(map[string]repoConfig)
	orgReference := make(map[string]orgConfig)

	for i := range config.Repo {
		owner := config.Repo[i].Owner
		name := config.Repo[i].Name
		hook, resp, err := githubClient.Repositories.CreateHook(ctx, owner, name, &githubHook)
		if resp == nil && err != nil {
			fmt.Printf("ERROR - %v repo hook placement: %v, response status: %v\n", *hook.Name, err, resp.Status)
		}
		repoReference[fmt.Sprint(owner, "/", name)] = config.Repo[i]
	}

	for j := range config.Org {
		org := config.Org[j].Org
		hook, resp, err := githubClient.Organizations.CreateHook(ctx, org, &githubHook)
		if resp == nil && err != nil {
			fmt.Printf("ERROR - %v org hook placement: %v, response status: %v\n", *hook.Name, err, resp.Status)
		}
		orgReference[org] = config.Org[j]
	}

	// slackClient := slack.New(config.sSlack)

	/*
	   hookHandler := func(w http.ResponseWriter, r *http.Request) {
	       // 	// NOTE: I'm not sure this is how to reference that part of the JSON message.
	   	if repoName := r.FormValue("repository.name"); repoName != "" {
	           // repoOwner := r.FormValue("repository.owner.login")
	           // stuff goes here
	       } else if orgName := r.FormValue("organization.name"); orgName != "" {
	           // stuff goes here
	       }
	   	// action := r.FormValue("action")
	   	// url := r.FormValue("project_card.url")

	   }
	*/

	// 	// channels := reference[repo].Slack.Channels
	//
	// 	params := slack.PostMessageParameters{}
	//
	// 	// for i := range channels {
	// 	// 	slackClient.PostMessage(channels[i])
	// 	// }
	//
	// })
	//
	// http.HandleFunc("/gpsihook", hookHandler)
	// http.ListenAndServe(":7000", nil)

	// Outline:
	// [X] - read from gpsi.toml file
	// [X] create github client
	// [X] create slack client
	// [X] place hooks on target repositories
	// [X] if hook exists
	// [X] - report it exists and move to next hook
	// [ ] create listening server
	// [ ] - start listenandserve
	// [ ] create hook handler
	// [ ] - listen for specific event lists
	// [ ] - create new slack message
	// [ ] - post to target channel(s)

	// Resources:
	// https://github.com/google/go-github/blob/master/github/projects.go
	// https://github.com/google/go-github/blob/master/github/repos_projects.go
	// https://github.com/google/go-github/blob/master/github/orgs_projects.go
}
