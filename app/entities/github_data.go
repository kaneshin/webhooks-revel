package entities

import (
	"encoding/json"
)

type GitHubUser struct {
	Login string `json:"login"`
	Id int `json:"id"`
	AvatarUrl string `json:"avatar_url"`
}

type GitHubIssueData struct {
	HtmlUrl string `json:"html_url"`
	Number int `json:"number"`
	State string `json:"state"`
	Title string `json:"title"`
	Body string `json:"body"`
	User GitHubUser `json:"user"`
}

type GitHubPullRequestData struct {
	HtmlUrl string `json:"html_url"`
	Number int `json:"number"`
	State string `json:"state"`
	Title string `json:"title"`
	Merged bool `json:"merged"`
	User GitHubUser `json:"user"`
}

type GitHubCommentData struct {
	HtmlUrl string `json:"html_url"`
	Body string `json:"body"`
	User GitHubUser `json:"user"`
}

type GitHubRepositoryData struct {
	Id int `json:"id"`
	Name string `json:"name"`
	FullName string `json:"full_name"`
	HtmlUrl string `json:"html_url"`
	Owner GitHubUser `json:"owner"`
}

// === IssueEvent

type GitHubIssueEventData struct {
	Action string `json:"action"`
	Issue GitHubIssueData `json:"issue"`
	Comment GitHubCommentData `json:"comment"`
	Repository GitHubRepositoryData `json:"repository"`
}

func NewGitHubIssueEventData(payload map[string]interface{}) *GitHubIssueEventData {
	d, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	data := &GitHubIssueEventData{}
	_ = json.Unmarshal(d, &data)
	return data
}

// === PullRequestEvent

type GitHubPullRequestEventData struct {
	Action string `json:"action"`
	Number int `json:"number"`
	PullRequest GitHubPullRequestData `json:"pull_request"`
	Comment GitHubCommentData `json:"comment"`
	Repository GitHubRepositoryData `json:"repository"`
}

func NewGitHubPullRequestEventData(payload map[string]interface{}) *GitHubPullRequestEventData {
	d, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	data := &GitHubPullRequestEventData{}
	_ = json.Unmarshal(d, &data)
	return data
}

