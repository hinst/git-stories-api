package git_stories_api

import (
	"time"
)

type DetailedLogEntryRow struct {
	Time       time.Time
	Parents    []ParentInfoEntry
	CommitHash string
}

type DetailedLogEntryRows []DetailedLogEntryRow

type ParentInfoEntry struct {
	CommitHash  string
	DiffSummary []DiffSummaryRow
}

func (rows DetailedLogEntryRows) FindByCommitHash(commitHash string) *DetailedLogEntryRow {
	for _, row := range rows {
		if row.CommitHash == commitHash {
			return &row
		}
	}
	return nil
}
