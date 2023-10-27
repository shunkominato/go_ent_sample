package slackNotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Slack struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconURL   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

func ErrorNotificator(userId string, errorLocation string, errMessage interface{}, value interface{}) {
	message := fmt.Sprintf(`
user_id: %v
error_location: %v
err_message: %v
%v`, userId, errorLocation, errMessage, value)

	params := Slack{
			Text:      message,
			Username:  "error",
			IconEmoji: ":gopher:",
			IconURL:   "",
			Channel:   os.Getenv("SLACK_CHANNEL"),
	}
	jsonparams, _ := json.Marshal(params)

	resp, _ := http.PostForm(
		os.Getenv("SLACK_WEBHOOK_URL"),
		url.Values{"payload": {string(jsonparams)}},
	)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
}