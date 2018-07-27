package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test_handler(t *testing.T) {
	c = config{
		SlackWebhook:        "test-webhook",
		GitHubWebhookSecret: "test-secret",
		GitHubRepoOwner:     "test-owner",
		GitHubRepoName:      "test-repo",
	}

	tests := []struct {
		desc string
		sec  string
		req  events.APIGatewayProxyRequest
		pass bool
	}{
		{
			"incorrect signature provided",
			"wrong-secret",
			events.APIGatewayProxyRequest{
				Headers: make(map[string]string),
				Body:    "test body",
			},
			false,
		},
		{
			"incorrect event provided",
			c.GitHubWebhookSecret,
			events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"X-GitHub-Event": "wrong-event",
				},
				Body: "test-body",
			},
			false,
		},
		{
			"no json body",
			c.GitHubWebhookSecret,
			events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"X-GitHub-Event": "project_card",
				},
			},
			false,
		},
	}

	for i, test := range tests {
		mac := hmac.New(sha1.New, []byte(test.sec))
		mac.Write([]byte(test.req.Body))
		sig := "sha1=" + hex.EncodeToString(mac.Sum(nil))
		test.req.Headers["X-Hub-Signature"] = sig

		_, err := handler(test.req)
		if (err == nil) != test.pass {
			t.Errorf("test #%d desc: %s, received error: %s", i+1, test.desc, err)
		}
	}

}
