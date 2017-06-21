package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

// SendSlackMessage ...
func SendSlackMessage(globalConfiguration *GlobalConfiguration, message string) {
	api := slack.New(globalConfiguration.Slack.Token)
	params := slack.PostMessageParameters{}

	attachment := &slack.Attachment{
		Fallback: message,
		Pretext:  "Push Dockerhub",
		Title:    message,
		Text:     ":blush: Ready",
	}

	params.Attachments = []slack.Attachment{*attachment}
	params.IconEmoji = ":dart:"

	channelID, timestamp, err := api.PostMessage(globalConfiguration.Slack.Channel, "", params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
