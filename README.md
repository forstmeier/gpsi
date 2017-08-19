# gpsi

### Description

<b>G</b>itHub <b>P</b>rojects <b>S</b>lack <b>I</b>ntegration (GPSI) -
pronounced "jip-see" - hooks into your
[GitHub Projects](https://help.github.com/articles/about-project-boards/)
board(s) and sends alerts to the [Slack](https://slack.com/) channel(s) of your
choice as your team moves issues cards across the kanban board.

## Status

So far, this will be a bare-bones skeleton with most of the code outline
available directly in the `gpsi.go` source file (given the project is a
relatively simple scope).

### Wishlist

This is a summary of several features that can be built / expanded upon from
the existing code (with as much description possible without becoming verbose).

- [ ] Selection from repos available to the authenticating user
  - [ ] e.g. Printout to the command line w/ selection options
- [ ] Command line configuration
  - [ ] Arguments following binary `gpsi` command (same as would be in TOML)
- [ ] Message details settings
  - [ ] Summary / details / additional alerts / etc.
- [ ] Kanban activity options
  - Card added / removed / moved / issue closed / etc.
