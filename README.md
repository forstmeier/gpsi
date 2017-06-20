# gpsi

---

### Description

GitHub Projects Slack Integration - pronounced "jip-see" - hooks into your
GitHub Projects board(s) and sends alerts to the Slack channel(s) of your
choice as your team moves issues cards across the kanban board.

Workflow:
- listens for GitHub Projects events (see go-github API)
- passes message into established Slack client / channel

Outline:
- usage options
  - fork repository, clone local copy, run "go build", execute binary
  - directly download available binary (release), execute binary
- execution options
  - gpsi.toml configuration, looks for file in GOPATH
  - flags, provided following binary call
- settings
  - GitHub
    - card added
    - card removed
    - cared moved
    - issue closed (maybe)
  - Slack
    - message details
    - message summary
    - channel options
    - additional messages / alerts
- errors
  - no project on repo (removed in settings)
  - invalid user token
  - repo does not exist
