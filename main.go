package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	webhooks "gopkg.in/go-playground/webhooks.v2"
	"gopkg.in/go-playground/webhooks.v2/github"

	gosxnotifier "github.com/deckarep/gosx-notifier"
)

const (
	path = "/payload"
	port = 4567
)

func main() {
	hook := github.New(&github.Config{Secret: os.Getenv("GITHUB_WEBHOOK_SECRET")})
	hook.RegisterEvents(HandleEvent, github.PullRequestEvent)

	err := webhooks.Run(hook, ":"+strconv.Itoa(port), path)

	if err != nil {
		log.Println(err)
	}
}

type message struct {
	title    string
	subtitle string
	body     string
	url      string
}

func newMessage(pr github.PullRequest, r github.Repository) message {
	return message{
		title:    r.Name,
		subtitle: pr.Title,
		body:     pr.Body,
		url:      pr.HTMLURL,
	}
}

// HandleEvent handles Github event
func HandleEvent(payload interface{}, header webhooks.Header) {
	switch payload.(type) {
	case github.PullRequestPayload:
		pullRequest := payload.(github.PullRequestPayload)
		if strings.Contains(pullRequest.PullRequest.Body, fmt.Sprintf("@%s", "tkbky")) {
			msg := newMessage(pullRequest.PullRequest, pullRequest.Repository)
			notify(msg)
		}
	}
}

func notify(message message) {
	note := gosxnotifier.NewNotification(message.body)
	note.Title = message.title
	note.Subtitle = message.subtitle
	note.Link = message.url
	note.AppIcon = "octocat.png"

	err := note.Push()

	if err != nil {
		log.Println(err)
	}
}
