package git_stories_api

import (
	"time"
)

type RepositoryLogEntry struct {
	Time       time.Time         `json:"time"`
	Parents    []ParentInfoEntry `json:"parents"`
	CommitHash string            `json:"commitHash"`
	AuthorName string            `json:"authorName"`
}

type RepositoryLogEntries []*RepositoryLogEntry

type ParentInfoEntry struct {
	CommitHash string           `json:"commitHash"`
	DiffRows   []DiffSummaryRow `json:"diffRows"`
}

func (rows RepositoryLogEntries) FindByCommitHash(commitHash string) *RepositoryLogEntry {
	for _, row := range rows {
		if row.CommitHash == commitHash {
			return row
		}
	}
	return nil
}
