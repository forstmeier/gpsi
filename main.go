package main

import (
	"encoding/json"
	"io/ioutil"

	slack "github.com/ashwanthkumar/slack-go-webhook"
)

type config struct {
	SlackWebhook string `json:"slack_webhook"`
}

func main() {
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	c := config{}
	if err = json.Unmarshal(content, &c); err != nil {
		panic(err)
	}

	payload := slack.Payload{
		Text: "Posting without a token!",
	}
	if err := slack.Send(c.SlackWebhook, "", payload); err != nil {
		panic(err)
	}
}
