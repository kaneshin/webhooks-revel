package models

var githubInstance *GitHub
func GetGitHubInstance() *GitHub {
	if githubInstance == nil {
		githubInstance = &GitHub{&WebHooks{}}
	}
	return githubInstance
}

var jiraInstance *Jira
func GetJiraInstance() *Jira {
	if jiraInstance == nil {
		jiraInstance = &Jira{&WebHooks{}}
	}
	return jiraInstance
}

var slackInstance *Slack
func GetSlackInstance() *Slack {
	if slackInstance == nil {
		slackInstance = &Slack{&WebHooks{}}
	}
	return slackInstance
}

