package infra

import (
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type Client struct {
	api    *slack.Client
	socket *socketmode.Client
}

func SlackInitialize(token string) (*Client, error) {
	api := slack.New(token, slack.OptionAppLevelToken(token))
	socket := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	return &Client{api: api, socket: socket}, nil
}
