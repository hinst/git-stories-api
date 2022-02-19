package git_stories_api

type StoriesRequest struct {
	LogEntries []*RepositoryLogEntry `json:"logEntries"`
	TimeZone   string                `json:"timeZone"`
}
