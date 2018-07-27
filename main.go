package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

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

func signBody(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}

func verifySignature(secret, body []byte, signature string) bool {
	const signaturePrefix = "sha1="
	const signatureLength = 45

	if len(signature) != signatureLength || !strings.HasPrefix(signature, signaturePrefix) {
		return false
	}

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature[5:]))

	return hmac.Equal(signBody(secret, body), actual)
}

func handler(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sig := r.Headers["X-Hub-Signature"]
	if !verifySignature([]byte(c.GitHubWebhookSecret), []byte(r.Body), sig) {
		return events.APIGatewayProxyResponse{}, errors.New("invalid github signature provided in request")
	}

	if evt, _ := r.Headers["X-GitHub-Event"]; evt != "project_card" {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("incorrect event type, expected project_card, received %s", evt)
	}

	raw := json.RawMessage(r.Body)
	bodyBytes, err := raw.MarshalJSON()
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("error creating raw message bytes, %s", err.Error())
	}

	var event github.ProjectCardEvent
	if err := json.Unmarshal(bodyBytes, &event); err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("error unmarshalling body, %s", err.Error())
	}

	owner, name := *event.Repo.Owner.Login, *event.Repo.Name
	if owner != c.GitHubRepoOwner || name != c.GitHubRepoName {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("incorrect repo, wanted %s/%s, received %s/%s", c.GitHubRepoOwner, c.GitHubRepoName, owner, name)
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
		return events.APIGatewayProxyResponse{}, fmt.Errorf("error sending slack message: %v", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
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
