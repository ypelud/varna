package main

// GlobalConfiguration ...
type GlobalConfiguration struct {
	Slack *slackInfo
}

// SlackParams ...
type slackInfo struct {
	Token   string
	Channel string
}
