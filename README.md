# Camel Spotting :camel: :telescope:

### Description

**Camel Spotting** is an integration that connects your [GitHub Projects](https://help.github.com/articles/about-project-boards/) updates into
your team's [Slack](https://slack.com/) to keep everyone on the same page. Be
in the know about your kanban board.

### Status

So far, this will be a bare-bones skeleton with most of the code outline
available directly in the `main.go` source file (given the project is
a relatively simple scope).

### Installation flow

This is very much up in the air and I'll be focusing on making this run both
locally and on AWS Lambda before going into things like the Slack button.

- Create app on Slack Apps site
  - Enable Incoming Webhooks
  - Copy webhook URL for chosen channel
- Place into config.json file:
  - GitHub repo name
  - GitHub project name (possibly)
  - GitHub webhook secret
  - Slack webhook URL
- Launch the app on AWS Lambda as a standalone function
- In the GitHub repo webhook Settings:
  - Set payload URL to the URL from AWS Lambda
  - Apply GitHub webhook secret
  - Select only "Project Cards" events
- App deployed to AWS Lambda:
  - Listens for new events on GitHub
  - Accepts/processes as needed
  - Posts message to specified webhook

### Wishlist / roadmap

This is a summary of several features that can be built / expanded upon from
the existing code (with as much description possible without becoming verbose).
NOTE: this is a "living list" so I'll add to stuff her periodically; I might
also just move this all over to GitHub Issues.

- [ ] streamline activation/installation
  - [ ] e.g. GitHub sign-in and Slack button
- [ ] configure setting through Slack slash commands (after install)
  - [ ] e.g. message verbosity, board actions, channels, etc.
