package git_stories_api

import "time"

type StoryEntry struct {
	Time           time.Time
	CommitHash     string
	ParentHash     string
	Description    string
	SourceFilePath string
}
