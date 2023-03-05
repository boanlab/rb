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
./rb -url=http://127.0.0.1:8080 -w=5 -r=1000 -t=500s

2023/03/05 18:28:35 Welcome to rb
2023/03/05 18:28:35 Running Benchmark with url=http://127.0.0.1:8080, workers=5, requestsPerWorker=1000, timeout=8m20s
Sent 5000 requests in 8m20s
Average request time: 18.6404 ms
Request time standard deviation: 15.7672 ms
Response Time Percentiles:
 10% in 10.192 ms
 25% in 11.130 ms
 50% in 12.737 ms
 75% in 18.205 ms
 90% in 37.208 ms
 95% in 41.972 ms
 99% in 101.438 ms
Status code statistics
200: 5000
```
## Features
- Simulates multiple concurrent requests to the API endpoint
- Allows users to evaluate the performance of the system under different levels of load
- Provides detailed statistics on response times and status codes
- Supports command line options to configure the test parameters


## Command Line Options

```
Options
Note that all workers perform tasks in parallel
If you want a synchronous test rather than sending multiple requests at once, do not create the -w option.
  -r int
        number of requests per worker (default 100)
  -t duration
        time out (default 1m0s)
  -url string
        target URL (default "http://localhost:8080")
  -w int
        number of workers (default 1)
```
> Requests are sent in parallel according to the specified number of workers. Each worker performs the requested number of requests individually. Therefore, if you don't want to send multiple requests at the same time, you can simply omit the -w option (default worker = 1)

## Development Roadmap
- Add support for additional HTTP methods (currently only GET is supported)
- Support for custom Request Body
- Additional features under development

## Contributors
[Younghwan Kim](https://github.com/royroyee)

## Contributing
This project is still under development and there are many areas that need improvement. Feedback and suggestions are always welcome through issues or pull requests

## License
[MIT License](https://github.com/boanlab/rb/blob/main/LICENSE)
