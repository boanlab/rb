package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"sync"
	"time"
)

func request(url string, workers int, requestsPerWorker int, timeout time.Duration) {

	httpClient := &fasthttp.Client{}

	var wg sync.WaitGroup
	wg.Add(workers)

	responses := make([][]response, workers) // slice of responses, one per worker

	for i := 0; i < workers; i++ {
		go func(workerNum int) {
			defer wg.Done()

			for j := 0; j < requestsPerWorker; j++ {
				req := fasthttp.AcquireRequest()
				req.SetRequestURI(url)

				req.Header.SetMethod("GET")
				resp := fasthttp.AcquireResponse()

				startTime := time.Now()
				err := httpClient.Do(req, resp)

				if err != nil {
					fmt.Println("Request failed: ", err)
				} else {
					responseTime := time.Since(startTime).Seconds() * 1000
					responses[workerNum] = append(responses[workerNum], response{time: responseTime, status: resp.StatusCode()})
				}

				fasthttp.ReleaseRequest(req)
				fasthttp.ReleaseResponse(resp) // avoid memory leak
			}
		}(i)
	}

	stop := time.NewTimer(timeout)

	go func() {
		<-stop.C
		rt := getResponsesTimes(responses)

		// print statistics
		printStatistics(rt, timeout)

		// print response time percentile
		printPercentiles(rt)

		// Print status code statistics
		printStatusCodes(getResponseStatuses(responses))
	}()

	wg.Wait()

	rt := getResponsesTimes(responses)

	// print statistics
	printStatistics(rt, timeout)

	// print response time percentile
	printPercentiles(rt)

	// Print status code statistics
	printStatusCodes(getResponseStatuses(responses))
}
