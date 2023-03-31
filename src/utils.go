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
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

// original code is from https://github.com/rakyll/hey
func exitWithError(err string) {
	if err != "" {
		fmt.Fprintln(os.Stderr, "Error:", err)
		fmt.Fprintln(os.Stderr)
	}
	flag.Usage()
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

// helper function to flatten the responses slice of slice into a single slice
func getResponsesTimes(responses [][]Response) []float64 {
	var responseTimes []float64
	for _, workerResponses := range responses {
		for _, resp := range workerResponses {
			responseTimes = append(responseTimes, resp.Time)
		}
	}
	return responseTimes
}

// helper function to flatten the responseCodes slice of slice into a single slice
func getResponseStatuses(responses [][]Response) []int {
	var responseStatuses []int
	for _, workerResponses := range responses {
		for _, resp := range workerResponses {
			responseStatuses = append(responseStatuses, resp.Status)
		}
	}
	return responseStatuses
}

func printAll(responses [][]Response, timeout time.Duration, start time.Time) {

	// total time
	duration := time.Since(start)
	fmt.Printf("total time: %.2f seconds\n", float64(duration)/float64(time.Second))

	rt := getResponsesTimes(responses)

	printStatistics(rt, timeout)

	// print response time histogram
	printHistogram(rt)

	// print response time percentile
	printPercentiles(rt)

	// print status code statistics
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

func printHistogram(requestTimes []float64) {
	numBins := 15
	n := len(requestTimes)

	// Sort the request times
	sort.Float64s(requestTimes)

	// Calculate the min and max
	min := requestTimes[0]
	max := requestTimes[n-1]

	// TODO This method is not the most optimal solution so I'm currently exploring more efficient alternatives
	// To prevent excessive distortion of the histogram caused by high values, the last two values are excluded from the bin boundaries.
	secondMax := requestTimes[n-2]
	thirdMax := requestTimes[n-3]

	// Calculate the bin size based on the min and max
	binSize := (thirdMax - min) / float64(numBins)

	// Calculate the bin boundaries
	binBoundaries := make([]float64, numBins+1)
	for i := 0; i <= numBins-2; i++ {
		binBoundaries[i] = min + float64(i)*binSize
	}

	// TODO This method is not the most optimal solution so I'm currently exploring more efficient alternatives
	// To prevent excessive distortion of the histogram caused by high values, the last two values are excluded from the bin boundaries.
	binBoundaries[numBins-1] = secondMax
	binBoundaries[numBins] = max

	// Create bins and count values in each bin
	bins := make([]int, numBins)
	binIndex := 0
	for _, rt := range requestTimes {
		if rt <= binBoundaries[binIndex+1] {
			bins[binIndex]++
		} else {
			binIndex++
			if binIndex >= numBins {
				binIndex = numBins - 1
			}
			bins[binIndex]++
		}
	}

	// Print histogram
	fmt.Println("Response time histogram:")
	for i, binCount := range bins {
		binStart := binBoundaries[i]
		binEnd := binBoundaries[i+1]

		// Normalize the number of blocks to a maximum of 50
		numBlocks := int(math.Round((float64(binCount) / float64(n)) * 50))

		blocks := strings.Repeat("▄", numBlocks)
		fmt.Printf("%.3f - %.3f ms [%d] %s\n", binStart, binEnd, binCount, blocks)
	}
}
