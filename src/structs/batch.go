package structs

import "time"

type Batch struct {
	BatchID   int
	Timestamp time.Time
	Links     map[string]string
}
