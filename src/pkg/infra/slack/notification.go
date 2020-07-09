package slack

import (
	"fmt"
	"os"

	"github.com/ashwanthkumar/slack-go-webhook"
)

const (
	CHANNEL  = "botmake"
	USERNAME = "GoBot"
)

// TestNotification Test通知用
func TestNotification() {
	field1 := slack.Field{Title: "Message", Value: "Hello World!!!!"}
	field2 := slack.Field{Title: "AnythingKey", Value: "AnythingValue"}

	attachment := slack.Attachment{}
	attachment.AddField(field1).AddField(field2)
	color := "good"
	attachment.Color = &color
	payload := slack.Payload{
		Username:    USERNAME,
		Channel:     CHANNEL,
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.Send(os.Getenv("SLACK_WEBHOOK_URL"), "", payload)
	if err != nil {
		fmt.Println(err)
	}
}
