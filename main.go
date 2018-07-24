package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	slack "github.com/ashwanthkumar/slack-go-webhook"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/github"
)

type config struct {
	SlackWebhook        string `json:"slack_webhook"`
	GitHubWebhookSecret string `json:"github_webhook_secret"`
	GitHubRepoOwner     string `json:"github_repo_owner"`
	GitHubRepoName      string `json:"github_repo_name"`
}

var c config

func handler(r events.APIGatewayProxyRequest) error {
	if _, ok := r.Headers["X-Hub-Signature"]; !ok {
		return errors.New("github secret not provided in request header")
	}

	if evt, _ := r.Headers["X-GitHub-Event"]; evt != "project_card" {
		return fmt.Errorf("incorrect event type, expected project_card, received %s", evt)
	}

	raw := json.RawMessage(r.Body)
	bodyBytes, err := raw.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error creating raw message bytes, %s", err.Error())
	}

	var event github.ProjectCardEvent
	if err := json.Unmarshal(bodyBytes, &event); err != nil {
		return fmt.Errorf("error unmarshalling body, %s", err.Error())
	}

	owner, name := *event.Repo.Owner.Login, *event.Repo.Name
	if owner != c.GitHubRepoOwner || name != c.GitHubRepoName {
		return fmt.Errorf("incorrect repo, wanted %s/%s, received %s/%s", c.GitHubRepoOwner, c.GitHubRepoName, owner, name)
	}

	title := fmt.Sprintf("%s %s a card in %s", *event.Sender.Login, *event.Action, *event.Repo.FullName)
	value := fmt.Sprintf("Visit the repo <%s|here>", *event.Repo.HTMLURL)

	attachment := slack.Attachment{}
	attachment.AddField(
		slack.Field{Title: title, Value: value},
	)
	payload := slack.Payload{
		Username:    *event.Sender.Login,
		Text:        "Repo project board update",
		Attachments: []slack.Attachment{attachment},
	}

	if err := slack.Send(c.SlackWebhook, "", payload); len(err) > 0 {
		return fmt.Errorf("error sending slack message: %v", err)
	}

	return nil
}

func main() {
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(content, &c); err != nil {
		panic(err)
	}

	lambda.Start(handler)
}
