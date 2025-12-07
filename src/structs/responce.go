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
	BatchID   int               `json:"batch_id"`  // 1
	Links     map[string]string `json:"links"`     // {"google.com": "available", ...}
	Timestamp string            `json:"timestamp"` // "2025-12-07T22:30:45.123456Z"
}
