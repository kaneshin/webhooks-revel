package models

import (
	"net/url"
	"net/http"
	"strings"
	"bytes"
	"encoding/json"
	"webhooks/app/config"
	"webhooks/app/entities"
)

type Slack struct {
	*WebHooks
}

func (self *Slack) SetPayload(form url.Values) error {
	payload := map[string]interface{}{}
	data := []byte(form["payload"][0])
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}
	self.payload = payload
	return nil
}

// === API

func (self *Slack) setSlackAPIUri(method string) string {
	var buffer bytes.Buffer
	buffer.WriteString("https://slack.com/api/")
	buffer.WriteString(method)
	return buffer.String()
}

func (self *Slack) SendChatMessage(data *entities.SlackData) error {
	uri := self.setSlackAPIUri("chat.postMessage")
	payload := url.Values{}
	payload.Set("token", config.SlackAPIToken)
	payload.Set("channel", data.Channel)
	payload.Set("text", data.Text)
	payload.Set("icon_url", data.IconUrl)
	payload.Set("username", data.Username)
	payload.Set("link_names", "1")
	if _, err := http.PostForm(uri, payload); err != nil {
		return err
	}
	return nil
}

func WriteSlackLink(buffer *bytes.Buffer, link string, alias string) {
	buffer.WriteString("<")
	buffer.WriteString(link)
	buffer.WriteString("|")
	buffer.WriteString(alias)
	buffer.WriteString(">")
}


// === service

func (self *Slack) setSlackServiceUri(service string, val url.Values) string {
	var buffer bytes.Buffer
	buffer.WriteString("https://")
	buffer.WriteString(config.SlackTeamIdentifier)
	buffer.WriteString(".slack.com/services/hooks/")
	buffer.WriteString(service)
	buffer.WriteString("?")
	buffer.WriteString(val.Encode())
	return buffer.String()
}

func (self *Slack) SendIncoming(data *entities.SlackData) error {
	val := url.Values{"token": {config.SlackIncomingToken}}
	uri := self.setSlackServiceUri("incoming-webhook", val)
	b, err := json.Marshal(data);
	if err != nil {
		return err
	}
	payload := url.Values{"payload": {string(b[:])}}
	if _, err = http.PostForm(uri, payload); err != nil {
		return err
	}
	return nil
}

func (self *Slack) SendBot(data *entities.SlackData) error {
	val := url.Values{"token": {config.SlackBotToken}}
	val.Set("channel", data.Channel)
	uri := self.setSlackServiceUri("slackbot", val)
	buf := strings.NewReader(data.Text)
	if _, err := http.Post(uri, "text/plain", buf); err != nil {
		return err
	}
	return nil
}
