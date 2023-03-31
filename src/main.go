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
	"net/http"
	"os"
	"runtime"
	"time"
)

var usage = `Usage rb : https://github.com/boanlab/rb  

[Currently, rb does not support HTTP/3 and only supports the GET method]

rb Options:

  -url : request single url (default "http://localhost:8080")

  -r :  number of total requests (default 100)
	  ex) -r=1000

  -t : time out (default 1m0s)
      ex) -60s

  -w : number of workers (default 10)
	  ex) -w=100

  -type : type of request
	f : ensures that the total number of requests is evenly divisible by the number of workers
	s : keeps the session alive and repeats assigned requests until the timeout is reached	
	ex) -type=f

  -disable-keepalive : Prevents reuse of TCP connections  (default false)
	ex) -disable-keepalive=true

---------------------------------------------------------------------------------------------------
  cpus : number of used cpu cores.
      (default for current machine is %d cores)

`

// Parse command-line arguments
var (
	url          = flag.String("url", "http://localhost:8080", "target URL")
	workers      = flag.Int("w", 10, "number of workers")
	totalRequest = flag.Int("r", 100, "number of requests per worker")
	timeout      = flag.Duration("t", 60*time.Second, "time out")
	requestType  = flag.String("type", "f", "type of request (-f / -s")

	disableKeepAlives = flag.Bool("disable-keepalive", false, "choose whether or not to enable the keep-alive feature")
)

func main() {

	// cpu core maximum
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, runtime.NumCPU()))
	}

	flag.Parse() // Parse command-line arguments

	fmt.Println()
	fmt.Println("Welcome to rb")

	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		exitWithError(err.Error())
	}

	c := &Request{
		requestType:       *requestType,
		HttpRequest:       req,
		Workers:           *workers,
		TotalRequests:     *totalRequest,
		Timeout:           *timeout,
		DisableKeepAlives: *disableKeepAlives,
	}

	fmt.Printf("Running Benchmark with type=%s url=%s, workers=%d, total requests =%d, timeout=%s", c.requestType, *url, c.Workers, c.TotalRequests, c.Timeout)
	c.runRequest()
	fmt.Println("-------------------------------------------------------------------------------------------------------")

}
