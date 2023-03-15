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
./rb -url=http://127.0.0.1:8080  -type=f -w=100 -r=3000 -t=100s

Welcome to rb
Running Benchmark with type=requestFixedPerWorker url=http://127.0.0.1:8080, workers=100, total requests =3000, timeout=1m40s

Sent 3000 requests in 1m40s
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
200: 3000

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
        c : sends concurrent requests from the specified number of workers within the given timeout
        
        ex) -type f
---------------------------------------------------------------------------------------------------
  cpus : number of used cpu cores.
      (default for current machine is 8 cores)


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
