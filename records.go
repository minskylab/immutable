package immutable

import (
	"strings"
	"time"

	"github.com/pkg/errors"
)

type ImmutableRecord struct {
	CID  string
	Date time.Time
}

// RecordLine returns a string representation of the record.
func (ir *ImmutableRecord) RecordLine() string {
	return ir.Date.Format(time.RFC3339Nano) + " " + ir.CID
}

func parseFromLine(line string) (ImmutableRecord, error) {
	parts := strings.Split(line, " ")

	date, err := time.Parse(time.RFC3339Nano, parts[0])
	if err != nil {
		return ImmutableRecord{}, errors.WithStack(err)
	}

	return ImmutableRecord{
		CID:  parts[1],
		Date: date,
	}, nil
}
