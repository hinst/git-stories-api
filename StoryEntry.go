package git_stories_api

import "time"

type StoryEntryFileChange struct {
	SourceFilePath string `json:"sourceFilePath"`
	Description    string `json:"description"`
}

type StoryEntryChangeset struct {
	Time        time.Time              `json:"time"`
	CommitHash  string                 `json:"commitHash"`
	FileEntries []StoryEntryFileChange `json:"fileEntries"`
}
