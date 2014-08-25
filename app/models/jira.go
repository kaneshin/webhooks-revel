package models

type Jira struct {
	*WebHooks
}

func NewJiraInstance() *Jira {
	jiraInstance = &Jira{&WebHooks{}}
	return jiraInstance
}

