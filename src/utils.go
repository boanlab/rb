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

// helper function to count the status codes in the given slice
func printStatusCodes(codes []int) {
	statusCodes := make(map[int]int)
	for _, code := range codes {
		statusCodes[code]++
	}
	fmt.Println("Status code statistics")
	for code, count := range statusCodes {
		fmt.Printf("%d: %d\n", code, count)
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
