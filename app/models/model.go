package models

var githubInstance *GitHub

func GetGitHubInstance() *GitHub {
	if githubInstance == nil {
		githubInstance = NewGitHubInstance()
	}
	return githubInstance
}

var slackInstance *Slack

func GetSlackInstance() *Slack {
	if slackInstance == nil {
		slackInstance = NewSlackInstance()
	}
	return slackInstance
}

var jiraInstance *Jira

func GetJiraInstance() *Jira {
	if jiraInstance == nil {
		jiraInstance = NewJiraInstance()
	}
	return jiraInstance
}
