package git_stories_api

import (
	"fmt"
	"strconv"
)

const (
	DiffSummaryBinary int = -1
	DiffSummaryError  int = -2
)

type DiffSummaryRow struct {
	FilePath       string
	InsertionCount int
	DeletionCount  int
}

var _ fmt.Stringer = &DiffSummaryRow{}

func (row *DiffSummaryRow) String() string {
	return row.FilePath +
		" +" + strconv.Itoa(row.InsertionCount) + " -" + strconv.Itoa(row.DeletionCount)
}
