package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	url, timeout, requests, concurrency := ParseFlags()

	start := time.Now()

	resutls := Worker(url, timeout, requests, concurrency)

	completed, failed, avgLatency, fastest, slowest := GetStats(resutls)

	requestsPerSecond := float64(completed) / time.Since(start).Seconds()

	fmt.Println("Load Test Completed")
	fmt.Println("---- Results ----")
	fmt.Printf("Completed requests: %d\n", completed)
	fmt.Printf("Failed Requests: %d\n", failed)
	fmt.Printf("Average Latency: %vms\n", math.Round(avgLatency.Seconds()*1000))
	fmt.Printf("Fastest: %vms\n", math.Round(fastest.Seconds()*1000))
	fmt.Printf("Slowest: %vms\n", math.Round(slowest.Seconds()*1000))
	fmt.Printf("Requests per second: %v\n", math.Round(requestsPerSecond))
}
