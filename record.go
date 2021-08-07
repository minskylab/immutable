package immutable

import "time"

type ImmutableRecord struct {
	CID       string
	PinataPin string
	Date      time.Time
}

func (ir *ImmutableRecord) RecordLine() string {
	return ir.Date.String() + " " + ir.CID + " " + ir.PinataPin
}
