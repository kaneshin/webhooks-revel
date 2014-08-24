package controllers

import (
	"github.com/revel/revel"
	"webhooks/app/config"
	"webhooks/app/entities"
	"webhooks/app/models"
)

type GitHub struct {
	*revel.Controller
}

func (c GitHub) Index() revel.Result {
	return c.Render()
}

func (c GitHub) Slack() revel.Result {
	github := models.GetGitHubInstance()
	if err := github.SetPayload(c.Params.Form); err != nil {
		panic(err)
	}
	result := map[string]interface{}{}
	slack := models.GetSlackInstance()
	slackData := entities.NewSlackData()
	if github.PushEvent() {
	}
	if data := github.PullRequestEvent(); data != nil {
		text := github.GetPullRequestText(data)
		slackData.InitializeSlackData(config.SlackIncomingChannel, text)
		slackData.Username = data.PullRequest.User.Login
		slackData.IconUrl = data.PullRequest.User.AvatarUrl
	}
	if data := github.PullRequestReviewCommentEvent(); data != nil {
		text := github.GetPullRequestReviewCommentText(data)
		slackData.InitializeSlackData(config.SlackIncomingChannel, text)
		slackData.Username = data.Comment.User.Login
		slackData.IconUrl = data.Comment.User.AvatarUrl
	}
	if data := github.IssueCommentEvent(); data != nil {
		text := github.GetIssueCommentText(data)
		slackData.InitializeSlackData(config.SlackIncomingChannel, text)
		slackData.Username = data.Comment.User.Login
		slackData.IconUrl = data.Comment.User.AvatarUrl
	}
	status := "OK"
	if err := slack.SendChatMessage(slackData); err != nil {
		status = "NG"
	}
	result["status"] = status
	return c.RenderJson(result)
}

