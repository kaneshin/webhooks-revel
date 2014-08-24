package controllers

import "github.com/revel/revel"

type Slack struct {
	*revel.Controller
}

func (c Slack) Index() revel.Result {
	return c.Render()
}

