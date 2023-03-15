package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Println("Options")
	fmt.Println("Note that all workers perform tasks in parallel")
	fmt.Println("If you want a synchronous test rather than sending multiple requests at once, do not create the -w option.")

	flag.PrintDefaults()
}

func main() {

	// cpu core maximum
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Parse command-line arguments
	var (
		url          = flag.String("url", "http://localhost:8080", "target URL")
		workers      = flag.Int("w", 10, "number of workers")
		totalRequest = flag.Int("r", 100, "number of requests per worker")
		timeout      = flag.Duration("t", 60*time.Second, "time out")
		requestType  = flag.String("type", "f", "type of request (-f / -s / -c)")
	)

	// Assign custom usage message
	flag.Usage = usage

	flag.Parse()

	fmt.Println("Welcome to rb")

	switch *requestType {
	case "f":
		fmt.Printf("Running Benchmark with type=requestFixedPerWorker url=%s, workers=%d, total requests =%d, timeout=%s\n", *url, *workers, *totalRequest, *timeout)
		fmt.Println("-------------------------------------------------------------------------------------------------------")
		requestFixedPerWorker(*url, *workers, *totalRequest, *timeout)

	case "s":
		fmt.Printf("Running Benchmark with type=requestSustained url=%s, workers=%d, total requests =%d, timeout=%s\n", *url, *workers, *totalRequest, *timeout)
		fmt.Println("-------------------------------------------------------------------------------------------------------")
		requestSustained(*url, *workers, *totalRequest, *timeout)

	case "c":
		fmt.Printf("Running Benchmark with type=requestConcurrently url=%s, workers=%d, total requests =%d, timeout=%s\n", *url, *workers, *timeout)
		fmt.Println("-------------------------------------------------------------------------------------------------------")
		requestConcurrently(*url, *workers, *timeout)

	default:
		fmt.Println("Invalid request type. Please use 'f' or 's' or 'c'")
		flag.Usage()
		os.Exit(1)
	}

}
