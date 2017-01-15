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

// HandleEvent handles Github event
func HandleEvent(payload interface{}, header webhooks.Header) {
	switch payload.(type) {
	case github.PullRequestPayload:
		pullRequest := payload.(github.PullRequestPayload)
		if strings.Contains(pullRequest.PullRequest.Body, fmt.Sprintf("@%s", "tkbky")) {
			notify(pullRequest.Repository.FullName, pullRequest.PullRequest.Title, pullRequest.PullRequest.HTMLURL, pullRequest.PullRequest.Body)
		}
	}
}

func notify(title string, subtitle string, link string, msg string) {
	note := gosxnotifier.NewNotification(msg)
	note.Title = title
	note.Subtitle = subtitle
	note.Link = link

	err := note.Push()

	if err != nil {
		log.Println(err)
	}
}
