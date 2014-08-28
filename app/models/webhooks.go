package models

import (
	"net/url"
)

type Behavior interface {
	SetPayload(form url.Values) error
}

type WebHooks struct {
	payload map[string]interface{}
}

func (self *WebHooks) GetPayload() interface{} {
	return self.payload
}
