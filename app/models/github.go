package models

import (
	"net/url"
	"encoding/json"
	"strings"
	"bytes"
	"strconv"
	"webhooks/app/entities"
)

type GitHub struct {
	*WebHooks
}

func (self *GitHub) SetPayload(form url.Values) error {
	payload := map[string]interface{}{}
	data := []byte(form["payload"][0])
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}
	self.payload = payload
	return nil
}

/**
 * IssueCommentEvent
 *
 * https://developer.github.com/v3/activity/events/types/#issuecommentevent
 */
func (self *GitHub) IssueCommentEvent() *entities.GitHubIssueEventData {
	data := entities.NewGitHubIssueEventData(self.GetPayload().(map[string]interface{}))
	if &data.Issue != nil && &data.Comment != nil {
		if strings.ToLower(data.Action) == "created" {
			return data
		}
	}
	return nil
}

func (self *GitHub) GetIssueCommentText(data *entities.GitHubIssueEventData) string {
	var buffer bytes.Buffer
	WriteSlackLink(&buffer, data.Comment.HtmlUrl, "The comment")
	buffer.WriteString(" created on ")
	WriteSlackLink(&buffer, data.Issue.HtmlUrl, "GitHub#" + strconv.Itoa(data.Issue.Number))
	buffer.WriteString(" [")
	WriteSlackLink(&buffer, data.Repository.HtmlUrl, data.Repository.FullName)
	buffer.WriteString("]\n>")
	buffer.WriteString(data.Comment.Body)
	return buffer.String()
}

/**
 * PushEvent
 *
 * https://developer.github.com/v3/activity/events/types/#pushevent
 */
func (self *GitHub) PushEvent() bool {
	return self.payload["commits"] != nil
}

func (self *GitHub) GetPushText() string {
	var buffer bytes.Buffer
	buffer.WriteString("Pushed")
	return buffer.String()
}

/**
 * PullRequestEvent
 *
 * https://developer.github.com/v3/activity/events/types/#pullrequestevent

 * action:
 *      assigned, unassigned,
 *      labeled, unlabeled,
 *      opened, closed, reopened, synchronize
 *
 * action == closed and merged == false => the pull request was closed with unmerged commits
 * action == closed and merged == true => the pull request was merged.
 */
func (self *GitHub) PullRequestEvent() *entities.GitHubPullRequestEventData {
	data := entities.NewGitHubPullRequestEventData(self.GetPayload().(map[string]interface{}))
	if &data.PullRequest != nil {
		if strings.ToLower(data.Action) != "created" {
			return data
		}
	}
	return nil
}

func (self *GitHub) GetPullRequestText(data *entities.GitHubPullRequestEventData) string {
	var buffer bytes.Buffer
	buffer.WriteString(" created on ")
	WriteSlackLink(&buffer, data.PullRequest.HtmlUrl, "GitHub#" + strconv.Itoa(data.PullRequest.Number))
	buffer.WriteString(" [")
	WriteSlackLink(&buffer, data.Repository.HtmlUrl, data.Repository.FullName)
	buffer.WriteString("]\n>")
	return buffer.String()
}


/**
 * PullRequestReviewCommentEvent
 *
 * https://developer.github.com/v3/activity/events/types/#pullrequestreviewcommentevent
 */
func (self *GitHub) PullRequestReviewCommentEvent() *entities.GitHubPullRequestEventData {
	data := entities.NewGitHubPullRequestEventData(self.GetPayload().(map[string]interface{}))
	if &data.PullRequest != nil && &data.Comment != nil {
		if strings.ToLower(data.Action) == "created" {
			return data
		}
	}
	return nil
}

func (self *GitHub) GetPullRequestReviewCommentText(data *entities.GitHubPullRequestEventData) string {
	var buffer bytes.Buffer
	WriteSlackLink(&buffer, data.Comment.HtmlUrl, "The comment")
	buffer.WriteString(" created on ")
	WriteSlackLink(&buffer, data.PullRequest.HtmlUrl, "GitHub#" + strconv.Itoa(data.PullRequest.Number))
	buffer.WriteString(" [")
	WriteSlackLink(&buffer, data.Repository.HtmlUrl, data.Repository.FullName)
	buffer.WriteString("]\n>")
	buffer.WriteString(data.Comment.Body)
	return buffer.String()
}

