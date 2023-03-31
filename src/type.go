package main

import (
	"net/http"
	"time"
)

type Request struct {
	requestType string // request type of rb options

	HttpRequest       *http.Request
	Workers           int           // total number of concurrent workers to run
	TotalRequests     int           // the total number of requests
	Timeout           time.Duration // timeout(sec)
	DisableKeepAlives bool
}

// define struct to hold response time and status code
type Response struct {
	Time   float64
	Status int
}
