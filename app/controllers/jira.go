package controllers

import "github.com/revel/revel"

type Jira struct {
	*revel.Controller
}

func (c Jira) Index() revel.Result {
	return c.Render()
}

