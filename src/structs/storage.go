type StorageState struct {
	NextBatchID int            `json:"next_batch_id"`
	Batches     []StorageBatch `json:"batches"`
}

type StorageBatch struct {
	BatchID   int               `json:"batch_id"`
	Timestamp string            `json:"timestamp"`
	Links     map[string]string `json:"links"`
}