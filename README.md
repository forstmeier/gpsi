# gpsi

### Description

<b>G</b>itHub <b>P</b>rojects <b>S</b>lack <b>I</b>ntegration (GPSI) -
pronounced "jip-see" - hooks into your
[GitHub Projects](https://help.github.com/articles/about-project-boards/)
board(s) and sends alerts to the [Slack](https://slack.com/) channel(s) of your
choice as your team moves issues cards across the kanban board.

### Status

So far, this will be a bare-bones skeleton with most of the code outline
available directly in the `gpsi.go` source file (given the project is a
relatively simple scope).

### Configuration

To set the environment variables needed to run this service for your project,
create a `gpsi.toml` file and place it in your local `gpis/` directory. Fill
out the variables listed below.

Create a [GitHub token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/)
with the scopes `repo` and `admin:repo_hook`. NOTE: These scopes may be subject
to change (I just haven't evaluated the minimum scopes needed for Project
Board event notifications).

Create a [Slack token](https://get.slack.help/hc/en-us/articles/215770388-Create-and-regenerate-API-tokens);
I do not currently know what the minimum scopes needed for this are.

```
[github]
token = "github-token"
owner = "repo-owner"
repos = ["repo-name"]

[slack]
token = "slack-token"
user = "slack-user"
team = "slack-user-team"
channels = ["slack-user-channels"]
```

### Wishlist / roadmap

This is a summary of several features that can be built / expanded upon from
the existing code (with as much description possible without becoming verbose).

- [ ] Streamline token generate
  - [ ] Automate via signup authentication (vs physically generating)
- [ ] Selection from repos available to the authenticating user
  - [ ] e.g. Printout to the command line w/ selection options
- [ ] Command line configuration
  - [ ] Arguments following binary `gpsi` command (same as would be in TOML)
- [ ] Message details settings
  - [ ] Summary / details / additional alerts / etc.
- [ ] Kanban activity options (as a setting)
  - Card added / removed / moved / issue closed / etc.
