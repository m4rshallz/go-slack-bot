package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {

	godotenv.Load(".env")

	token := os.Getenv("SLACK_AUTH_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")

	// set proxy
	if err := os.Setenv("HTTP_PROXY", os.Getenv("HTTP_PROXY")); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("SET HTTP_PROXY DONE")
	}

	if err := os.Setenv("HTTPS_PROXY", os.Getenv("HTTPS_PROXY")); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("SET HTTPS_PROXY DONE")
	}

	titleMessage := os.Args[1]
	keyMsg := os.Args[2]

	client := slack.New(token, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Pretext: "Application Notification",
		Text:    titleMessage,
		Color:   "4af030",
		Fields: []slack.AttachmentField{
			{
				// Title: keyMsg,
				Value: time.Now().Local().Format("2006-01-02 15:04:05") + fmt.Sprintf(" - %s", keyMsg),
			},
		},
	}

	_, timestamp, err := client.PostMessage(
		channelID,

		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Message sent at %s", timestamp)
}
