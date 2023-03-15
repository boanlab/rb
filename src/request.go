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
	"fmt"
	"net/http"
	"sync"
	"time"
)

func requestFixedPerWorker(url string, workers int, totalRequest int, timeout time.Duration) {
	var wg sync.WaitGroup

	if totalRequest%workers != 0 {
		fmt.Println("Error: Total requests must be evenly divisible by the number of workers.")
		return
	}

	httpClient := &http.Client{}

	wg.Add(workers)

	responses := make([][]response, workers) // slice of responses, one per worker
	requestsPerWorker := totalRequest / workers

	for i := 0; i < workers; i++ {
		go func(workerNum int) {
			defer wg.Done()

			for j := 0; j < requestsPerWorker; j++ {
				sendRequest(httpClient, url, workerNum, responses)
			}
		}(i)
	}
	stop := time.NewTimer(timeout)

	go func() {
		<-stop.C
		rt := getResponsesTimes(responses)
		// print all statistics
		printAll(rt, timeout, responses)

	}()
	wg.Wait()

	rt := getResponsesTimes(responses)
	// print all statistics
	printAll(rt, timeout, responses)
}

func requestSustained(url string, workers int, totalRequest int, timeout time.Duration) {

	var wg sync.WaitGroup

	if totalRequest%workers != 0 {
		fmt.Println("Error: Total requests must be evenly divisible by the number of workers.")
		return
	}

	transport := &http.Transport{
		IdleConnTimeout:     timeout, // to reuse connection
		MaxIdleConnsPerHost: workers,
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	wg.Add(workers)

	responses := make([][]response, workers) // slice of responses, one per worker
	requestsPerWorker := totalRequest / workers

	for i := 0; i < workers; i++ {
		go func(workerNum int) {
			defer wg.Done()

			stop := time.NewTimer(timeout)

			for {
				select {
				case <-stop.C:
					return

				default:
					for j := 0; j < requestsPerWorker; j++ {
						success := sendRequest(httpClient, url, workerNum, responses)
						if !success {
							continue
						}
					}
				}
			}
		}(i)
	}
	stop := time.NewTimer(timeout)

	go func() {
		<-stop.C
		rt := getResponsesTimes(responses)
		// print all statistics
		printAll(rt, timeout, responses)

	}()
	wg.Wait()
}

func requestConcurrently(url string, workers int, timeout time.Duration) {
	var wg sync.WaitGroup

	httpClient := &http.Client{}

	wg.Add(workers)

	responses := make([][]response, workers) // slice of responses, one per worker

	for i := 0; i < workers; i++ {

		go func(workerNum int) {
			defer wg.Done()
			sendRequest(httpClient, url, workerNum, responses)
		}(i)
	}
	stop := time.NewTimer(timeout)

	go func() {
		<-stop.C
		rt := getResponsesTimes(responses)
		// print all statistics
		printAll(rt, timeout, responses)

	}()
	wg.Wait()

	rt := getResponsesTimes(responses)
	// print all statistics
	printAll(rt, timeout, responses)
}

// send Request to ENDPOINT
func sendRequest(httpClient *http.Client, url string, workerNum int, responses [][]response) (result bool) {
	req, err := http.NewRequest("GET", url, nil)

	startTime := time.Now().UnixNano()
	resp, err := httpClient.Do(req)
	endTime := time.Now().UnixNano()

	if err != nil {
		fmt.Println("Request failed: ", err)
		return false
	} else {
		responseTime := time.Duration(endTime-startTime) * time.Nanosecond
		responses[workerNum] = append(responses[workerNum], response{time: responseTime.Seconds() * 1000, status: resp.StatusCode})
		return false
	}
	resp.Body.Close()

	return true
}
