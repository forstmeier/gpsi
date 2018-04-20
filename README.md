# gpsi

### Description

<b>G</b>itHub <b>P</b>rojects <b>S</b>lack <b>I</b>ntegration (GPSI) -
pronounced like the word "gypsy" - hooks into your
[GitHub Projects](https://help.github.com/articles/about-project-boards/)
board(s) and sends alerts to the [Slack](https://slack.com/) channel(s) of your
choice as your team moves issues cards across the kanban board.

### Status

So far, this will be a bare-bones skeleton with most of the code outline
available directly in the `gpsi.go` source file (given the project is a
relatively simple scope).

### Configuration

To set the environment variables needed to run this service for your project,
create a `gpsi.toml` file and place it in your local `gpsi/` directory. Fill
out the variables listed below.

Create a [GitHub token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/)
with the scopes `repo` and `admin:repo_hook`. NOTE: These scopes may be subject
to change (I just haven't evaluated the minimum scopes needed for Project
Board event notifications).

Create a [Slack token](https://get.slack.help/hc/en-us/articles/215770388-Create-and-regenerate-API-tokens);
I do not currently know what the minimum scopes needed for this are.

**NOTE**: This structure assumes that there is just ONE authenticating user for
the GitHub and Slack clients - support for different authenticating user tokens
could be something built in the future but might not ultimately make sense
if the "streamline" authenticating gets built soon.

### Wishlist / roadmap

This is a summary of several features that can be built / expanded upon from
the existing code (with as much description possible without becoming verbose).
NOTE: this is a "living list" so I'll add to stuff her periodically; I might
also just move this all over to GitHub Issues.

- [ ] Streamline GitHub / Slack activation
  - [ ] Improve token generation
    - [ ] Automate via signup authentication (vs physically generating)
    - [ ] E.g. OAuth2 on GitHub approval
  - [ ] Automate Slack integration (e.g. Slack Button)
    - [ ] Improve channel selection process
- [ ] Remotely host service (e.g. Heroku)
- [ ] Dynamically select from repos available to the authenticating user
  - [ ] e.g. Printout to the command line w/ selection options
- [ ] In-Slack configuration
  - [ ] Communicate w/ in-channel gpsi bot
    - [ ] Adjust setting, add channels, etc.
- [ ] Message details settings
  - [ ] Summary / details / additional alerts / etc.
- [ ] Kanban activity options (as a setting)
  - Card added / removed / moved / issue closed / etc.
