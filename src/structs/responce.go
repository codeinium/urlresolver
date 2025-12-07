type HistoryResponse struct {
	Batches []HistoryBatch `json:"batches"`
}

type HistoryBatch struct {
	BatchID    int               `json:"batch_id"`
	Timestamp  string            `json:"timestamp"`
	LinksCount int               `json:"links_count"`
	Links      map[string]string `json:"links"`
}

type CheckLinksResponse struct {
	BatchID   int               `json:"batch_id"`
	Links     map[string]string `json:"links"`
	Timestamp string            `json:"timestamp"`
}
