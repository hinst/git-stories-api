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
	FilePath       string `json:"filePath"`
	InsertionCount int    `json:"insertionCount"`
	DeletionCount  int    `json:"deletionCount"`
}

var _ fmt.Stringer = &DiffSummaryRow{}

func (row *DiffSummaryRow) String() string {
	return row.FilePath +
		" +" + strconv.Itoa(row.InsertionCount) + " -" + strconv.Itoa(row.DeletionCount)
}
