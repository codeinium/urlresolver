package structs

import "sync"

type AppState struct {
	NextBatchID int
	Batches     map[int]*Batch
	mu          sync.RWMutex
}
