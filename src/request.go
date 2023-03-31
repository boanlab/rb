//MIT License
//
//Copyright (c) 2023 BoanLab
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func (c Request) runRequest() {

	if (c.TotalRequests % c.Workers) != 0 {
		fmt.Println("Error: Total requests must be evenly divisible by the number of workers.")
		return
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         c.HttpRequest.Host,
		},
		MaxIdleConnsPerHost: c.Workers,
		DisableKeepAlives:   c.DisableKeepAlives,
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   c.Timeout,
	}

	responses := make([][]Response, c.Workers) // slice of responses, one per worker

	if c.requestType == "f" {
		c.requestFixedPerWorker(httpClient, responses)
	} else { // requestType == "s"
		c.requestSustained(httpClient, responses)
	}
}

func (c Request) requestFixedPerWorker(httpClient *http.Client, responses [][]Response) {

	var wg sync.WaitGroup

	wg.Add(c.Workers)

	start := time.Now() // Purpose of measuring the total time taken to send a request.

	requestsPerWorker := c.TotalRequests / c.Workers

	for i := 0; i < c.Workers; i++ {
		go func(workerNum int) {
			defer wg.Done()

			for j := 0; j < requestsPerWorker; j++ {
				c.sendRequest(c.HttpRequest, httpClient, workerNum, responses)

			}
		}(i)
	}
	stop := time.NewTimer(c.Timeout)

	go func() {
		<-stop.C
		// print all statistics
		printAll(responses, c.Timeout, start)

	}()
	wg.Wait()

	// print all statistics
	printAll(responses, c.Timeout, start)
}

func (c Request) requestSustained(httpClient *http.Client, responses [][]Response) {

	var wg sync.WaitGroup

	wg.Add(c.Workers)

	start := time.Now() // Purpose of measuring the total time taken to send a request.

	requestsPerWorker := c.TotalRequests / c.Workers

	for i := 0; i < c.Workers; i++ {
		go func(workerNum int) {
			defer wg.Done()
			stop := time.NewTimer(c.Timeout)

			for {
				select {
				case <-stop.C:
					// print all statistics
					printAll(responses, c.Timeout, start)
					os.Exit(1)
					return

				default:
					for j := 0; j < requestsPerWorker; j++ {
						c.sendRequest(c.HttpRequest, httpClient, workerNum, responses)
					}
				}
			}
		}(i)
	}
	stop := time.NewTimer(c.Timeout)

	go func() {
		<-stop.C
		// print all statistics
		printAll(responses, c.Timeout, start)
		os.Exit(1)

	}()
	wg.Wait()
}

// send Request to ENDPOINT
func (c Request) sendRequest(req *http.Request, httpClient *http.Client, workerNum int, responses [][]Response) {

	startTime := time.Now().UnixNano()
	resp, err := httpClient.Do(req)
	endTime := time.Now().UnixNano()

	if err != nil {
		fmt.Println("Request failed: ", err)
	}
	responseTime := time.Duration(endTime-startTime) * time.Nanosecond
	responses[workerNum] = append(responses[workerNum], Response{Time: responseTime.Seconds() * 1000, Status: resp.StatusCode})

	io.Copy(io.Discard, resp.Body) // to read the response body and discard it directly without copying the data to disk.
	resp.Body.Close()
}
