package git_stories_api

import (
	"time"
)

type DetailedLogEntryRow struct {
	Time       time.Time
	Parents    []ParentInfoEntry
	CommitHash string
}

type ParentInfoEntry struct {
	CommitHash  string
	DiffSummary []DiffSummaryRow
}
