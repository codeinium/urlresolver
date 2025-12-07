package services

import (
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	not_availible = "not_availible"
	availible     = "availible"
)

func checkUrl(url string) string {
	if url == "" {
		return availible
	}

	if !strings.HasPrefix(url, "http://") &&
		!strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Head(url)

	if err != nil {
		log.Printf("Ошибка при проверке %s: %v\n", url, err)
		return not_availible
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return availible
	}
	return not_availible
}

func CheckMuliplyLinks(urls []string) map[string]string {
	results := make(map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {

			defer wg.Done()
			status := checkUrl(u)
			mu.Lock()
			results[u] = status
			mu.Unlock()
		}(url)
	}

	wg.Wait()

	return results
}
