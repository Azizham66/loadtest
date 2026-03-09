# Loadtest

A minimal `wrk` clone written in Go for learning purposes. This is a simple HTTP load testing tool that demonstrates concurrent programming and performance measurement.

## Features

- HTTP load testing with configurable concurrency
- Configurable request count and timeout
- Real-time statistics (latency, throughput, success/failure rates)
- Simple command-line interface
- Safe concurrency with semaphore pattern

## Installation

```bash
git clone <repository-url>
cd loadtest
go build -o loadtest .
```

## Usage

```bash
./loadtest -url <URL> [options]
```

### Options

- `-url string`: URL to stress test (required)
- `-requests int`: Number of requests to send (default: 10)
- `-concurrency int`: Number of concurrent requests (default: 10)
- `-timeout int`: Timeout in seconds (default: 5)

### Examples

```bash
# Basic test with 100 requests
./loadtest -url=https://example.com -requests=100

# High concurrency test
./loadtest -url=https://api.example.com -requests=1000 -concurrency=50

# Test with custom timeout
./loadtest -url=https://slow-server.com -requests=500 -timeout=10
```

## Example Output

```
Load Test Completed
---- Results ----
Completed requests: 500
Failed Requests: 0
Average Latency: 85ms
Fastest: 71ms
Slowest: 272ms
Requests per second: 150
```

## Learning Goals

This project demonstrates:
- Go goroutines and channels
- Semaphore pattern for concurrency control
- HTTP client configuration
- Performance measurement and statistics
- Command-line argument parsing
- Error handling and resource cleanup

## Technical Details

- **Concurrency**: Limited using buffered channels as semaphores
- **Safety**: Proper goroutine cleanup with `sync.WaitGroup`
- **Accuracy**: Precise timing measurements for latency and throughput
- **Resource Management**: Automatic response body cleanup

## Limitations

- HTTP GET requests only
- No custom headers or request bodies
- Basic statistics (no percentiles)
- Single URL testing only

## Contributing

This is a learning project. Feel free to suggest improvements or extend functionality!

## License

MIT License - feel free to use for learning purposes.
