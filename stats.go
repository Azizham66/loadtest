package main

import "time"

// Completed requests, failed requests, avg latency, fastest, slowest, reqs/second

func GetStats(results <-chan Result) (int, int, time.Duration, time.Duration, time.Duration) {
	var completed int
	var failed int
	var avgLatency time.Duration
	var fastest time.Duration
	var slowest time.Duration
	var totalLatency time.Duration

	first := true
	totalResults := 0

	for result := range results {
		totalResults++
		completed++
		totalLatency += result.Latency
		if first {
			fastest = result.Latency
			slowest = result.Latency
			first = false
		}

		if result.Latency < fastest {
			fastest = result.Latency
		}

		if result.Latency > slowest {
			slowest = result.Latency
		}

		if !result.Success {
			failed++
		}
	}

	avgLatency = totalLatency / time.Duration(totalResults)
	return completed, failed, avgLatency, fastest, slowest
}
