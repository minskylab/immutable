package immutable

import "time"

type ImmutableRecord struct {
	CID       string
	PinataPin string
	Date      time.Time
}

func (ir *ImmutableRecord) RecordLine() string {
	return ir.Date.Format(time.RFC3339Nano) + " " + ir.CID // + " " + ir.PinataPin
}
