# rb
**A Simple REST API Benchmarking tool**

## Description
This tool enables users to measure the speed and responsiveness of REST API requests, and conduct load testing on web servers. This is achieved by simulating multiple concurrent requests to the API endpoint, allowing users to evaluate the performance of the system under different levels of load.

## Installation
```
git clone https://github.com/boanlab/rb.git
cd rb/src
go build
```

## Usage
```
$ ./rb -url=http://127.0.0.1:8080  -type=f -w=10 -r=1000 -t=60s

Welcome to rb

Running Benchmark with type=f url=http://serverhwan.shop:31111, workers=10, total requests =1000, timeout=1m0stotal 

time: 1.64 seconds
Sent 1000 requests in 1m0s
Average request time: 15.6467 ms
Request time standard deviation: 4.1981 ms
Response time histogram:
8.991 - 10.751 ms [35] ▄▄
10.751 - 12.511 ms [185] ▄▄▄▄▄▄▄▄▄
12.511 - 14.271 ms [234] ▄▄▄▄▄▄▄▄▄▄▄▄
14.271 - 16.030 ms [188] ▄▄▄▄▄▄▄▄▄
16.030 - 17.790 ms [127] ▄▄▄▄▄▄
17.790 - 19.550 ms [71] ▄▄▄▄
19.550 - 21.310 ms [73] ▄▄▄▄
21.310 - 23.070 ms [44] ▄▄
23.070 - 24.830 ms [11] ▄
24.830 - 26.590 ms [8] 
26.590 - 28.350 ms [3] 
28.350 - 30.109 ms [8] 
30.109 - 31.869 ms [3] 
31.869 - 35.552 ms [9] 
35.552 - 35.758 ms [1] 
Response Time Percentiles:
 10% in 11.416 ms
 25% in 12.799 ms
 50% in 14.646 ms
 75% in 17.366 ms
 90% in 20.948 ms
 95% in 22.508 ms
 99% in 31.303 ms
Status code statistics
200  :  1000

```

## Features
- Simulates multiple concurrent requests to the API endpoint
- Allows users to evaluate the performance of the system under different levels of load
- Provides detailed statistics on response times and status codes
- Supports command line options to configure the test parameters


## Command Line Options

```
Usage rb : https://github.com/boanlab/rb  
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


```
> Requests are sent in parallel according to the specified number of workers. Each worker performs the requested number of requests individually. Therefore, if you don't want to send multiple requests at the same time, you can simply use the -w=1



## Contributors
[Younghwan Kim](https://github.com/royroyee)

## Contributing
This project is still under development and there are many areas that need improvement. Feedback and suggestions are always welcome through issues or pull requests

## License
[MIT License](https://github.com/boanlab/rb/blob/main/LICENSE)
