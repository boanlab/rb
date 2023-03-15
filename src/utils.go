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
	"math"

	"sort"
	"time"
)

// define struct to hold response time and status code
type response struct {
	time   float64
	status int
}

// helper function to flatten the responses slice of slice into a single slice
func getResponsesTimes(responses [][]response) []float64 {
	var responseTimes []float64
	for _, workerResponses := range responses {
		for _, resp := range workerResponses {
			responseTimes = append(responseTimes, resp.time)
		}
	}
	return responseTimes
}

// helper function to flatten the responseCodes slice of slice into a single slice
func getResponseStatuses(responses [][]response) []int {
	var responseStatuses []int
	for _, workerResponses := range responses {
		for _, resp := range workerResponses {
			responseStatuses = append(responseStatuses, resp.status)
		}
	}
	return responseStatuses
}

func printAll(rt []float64, timeout time.Duration, responses [][]response) {
	printStatistics(rt, timeout)

	// print response time percentile
	printPercentiles(rt)

	// Print status code statistics
	printStatusCodes(getResponseStatuses(responses))

}

// helper function to count the status codes in the given slice
func printStatusCodes(codes []int) {
	statusCodes := make(map[int]int)
	for _, code := range codes {
		statusCodes[code]++
	}
	fmt.Println("Status code statistics")
	for code, count := range statusCodes {
		fmt.Println(code, " : ", count)
	}
}

// print statistics
func printStatistics(requestTimes []float64, timeout time.Duration) {
	requestCount := len(requestTimes)

	// calculate average request time
	var totalRequestTime float64
	for _, rt := range requestTimes {
		totalRequestTime += rt
	}
	avgRequestTime := totalRequestTime / float64(requestCount)

	// calculate request time standard deviation
	var rtVariance float64
	for _, rt := range requestTimes {
		rtVariance += (rt - avgRequestTime) * (rt - avgRequestTime)
	}
	rtStdDev := fmt.Sprintf("%.4f", math.Sqrt(rtVariance/float64(requestCount)))

	fmt.Println("Sent", requestCount, "requests in", timeout)
	fmt.Println("Average request time:", fmt.Sprintf("%.4f", avgRequestTime), "ms")
	fmt.Println("Request time standard deviation:", rtStdDev, "ms")

}

// print percentile statistics
func printPercentiles(data []float64) {
	percentiles := []float64{10, 25, 50, 75, 90, 95, 99}
	fmt.Println("Response Time Percentiles:")
	for _, percentile := range percentiles {
		value := getPercentile(data, percentile)
		fmt.Printf("%3.0f%% in %.3f ms\n", percentile, value)
	}
}

// calculate percentile for the given percentile range (0 to 100)
func getPercentile(data []float64, percentile float64) float64 {
	if len(data) == 0 {
		return 0
	}

	sort.Float64s(data)
	index := int(math.Ceil((percentile / 100) * float64(len(data))))
	return data[index-1]

}
