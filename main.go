package main

import (
	"bufio"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	b := true
	scanner := bufio.NewScanner(os.Stdin)

	var message []string
	for b {
		b = scanner.Scan()
		t := scanner.Text()
		message = append(message, t)
	}
	sendingMessage := strings.Join(message, "\n")

	v := url.Values{
		"token": {os.Getenv("OAuthAccessToken")},
		"channel": {os.Getenv("SlackChannelName")},
		"text": {sendingMessage},
	}

	http.PostForm("https://slack.com/api/chat.postMessage", v)
}