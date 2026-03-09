package main

import (
	"net/http"
	"sync"
	"time"
)

// Worker function

type Result struct {
	Success bool
	Latency time.Duration
}

const maxConcurrency = 100

// url: The URL to test
// timeout: The timeout for each request
// requests: The number of requests to send
// concurrency: The number of concurrent requests
func Worker(url string, timeout time.Duration, requests int, concurrency int) chan Result {
	var wg sync.WaitGroup
	sem := make(chan struct{}, min(concurrency, maxConcurrency))

	client := &http.Client{
		Timeout: timeout,
	}
	results := make(chan Result, requests)
	for range requests {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			start := time.Now()

			resp, err := client.Get(url)
			if resp != nil {
				defer resp.Body.Close()
			}

			latency := time.Since(start)

			if err != nil || resp.StatusCode >= 400 {
				results <- Result{Success: false, Latency: latency}
			} else {
				results <- Result{Success: true, Latency: latency}
			}

		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}
