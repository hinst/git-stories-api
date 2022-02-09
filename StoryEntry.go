package git_stories_api

import "time"

type StoryEntryChangeset struct {
	Time        time.Time
	CommitHash  string // Source commit hash
	FileEntries []StoryEntryFileChange
}

type StoryEntryFileChange struct {
	SourceFilePath string
	Description    string
}
