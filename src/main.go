package main

import (
	"flag"
	"fmt"
	"log"
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
	url := flag.String("url", "http://localhost:8080", "target URL")
	workers := flag.Int("w", 1, "number of workers")
	requestsPerWorker := flag.Int("r", 100, "number of requests per worker")
	timeout := flag.Duration("t", 60*time.Second, "time out")

	// Assign custom usage message
	flag.Usage = usage

	flag.Parse()

	log.Println("Welcome to rb")

	log.Printf("Running Benchmark with url=%s, workers=%d, requestsPerWorker=%d, timeout=%s\n", *url, *workers, *requestsPerWorker, *timeout)
	request(*url, *workers, *requestsPerWorker, *timeout)

}
