package git_stories_api

import (
	"time"
)

type RepositoryLogEntry struct {
	Time       time.Time
	Parents    []ParentInfoEntry
	CommitHash string
	AuthorName string
}

type RepositoryLogEntries []RepositoryLogEntry

type ParentInfoEntry struct {
	CommitHash  string
	DiffSummary []DiffSummaryRow
}

func (rows RepositoryLogEntries) FindByCommitHash(commitHash string) *RepositoryLogEntry {
	for _, row := range rows {
		if row.CommitHash == commitHash {
			return &row
		}
	}
	return nil
}
