package main

import (
	"flag"
	"log"
	"time"
)

func ParseFlags() (string, time.Duration, int, int) {
	var url string
	var timeoutInt int
	var requests int
	var concurrency int

	flag.StringVar(&url, "url", "", "URL to stress test")
	flag.IntVar(&timeoutInt, "timeout", 5, "Timeout in seconds")
	flag.IntVar(&requests, "requests", 10, "Number of requests to send")
	flag.IntVar(&concurrency, "concurrency", 10, "Number of concurrent requests")

	flag.Parse()

	if url == "" {
		log.Fatal("url flag is required, see -h or --help for help")
	}

	return url, time.Duration(timeoutInt) * time.Second, requests, concurrency
}
