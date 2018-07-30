# Camel Spotting :camel: :telescope:

### Description :octocat:

**[Camel Spotting](https://www.youtube.com/watch?v=6RexQLrcqwc)** is a tool that connects your [GitHub Projects](https://help.github.com/articles/about-project-boards/) updates into your team's [Slack](https://slack.com/) to keep everyone on the same page. Be in the know about your kanban board.

### Status :traffic_light:

Right now, working code is available in its entirety in the `main.go` file but it is still a work-in-progress.

### Installation flow :floppy_disk:

This is very much a work-in-progress I'll be focusing on making this run both locally and on AWS Lambda before going into things like the Slack button.

Add a file named `config.json` with the structure below to your local repository and then follow the steps below.

```json
{
  "slack_webhook": "webhook_url_value",
  "github_webhook_secret": "secret_value",
  "github_repo_owner": "repo_owner_value",
  "github_repo_name": "repo_name_value"
}
```

1. Create app on [Slack Apps site](https://api.slack.com/slack-apps)
    - Here you'll need to enable Incoming Webhooks and copy the URL for the
  channel you'd like to send message into.
2. Fill out the necessary fields in the `config.json` file.
    - Paste the copied Slack channel URL into the `slack_webhook` field.
    - Place your chosen GitHub webhook secret value into the
  `github_webhook_secret` field.
    - Add the target GitHub repository owner and name into the
  `github_repo_owner` and `github_repo_name` fields, respectively.
3. Create an AWS Lambda function that has an API Gateway trigger.
    - This can be done through the console page for AWS Lambda.
4. Bundle your application and user settings.
    - In your local environment, build the application binary (e.g.
  `go build -o /tmp/camelspotting`).
    - Compress both the binary file and the `config.json` file using zip (e.g.
  `zip -j camelspotting.zip camelspotting config.json`).
5. Set up your GitHub repository [webhook](https://developer.github.com/webhooks/creating/).
    - Set the **Payload URL** value to the URL for your AWS Lambda function
  API Gateway URL.
    - Paste your GitHub webhook secret from the `config.json` file into the
  **Secret** field.
    - Select "Let me select individual events" and make sure that only the
  **Project cards** is checked.
6. Upload the .zip file to AWS Lambda.
    - Using the console, upload your compressed file and you're ready to go!

### Wishlist / roadmap :bookmark:

This is a summary of several features that can be built / expanded upon from the existing code (with as much description possible without becoming verbose). NOTE: this is a "living list" so I'll add to stuff her periodically; I might also just move this all over to GitHub Issues as present them as future features.

- [ ] streamline activation/installation
  - [ ] e.g. GitHub sign-in and Slack button
- [ ] configure setting through Slack slash commands (after install)
  - [ ] e.g. message verbosity, board actions, channels, etc.
