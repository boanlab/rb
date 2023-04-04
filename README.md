# rb
**A Simple REST API Benchmarking tool**

## Description
This tool enables users to measure the speed and responsiveness of REST API requests, and conduct load testing on web servers. This is achieved by simulating multiple concurrent requests to the API endpoint, allowing users to evaluate the performance of the system under different levels of load.

## Installation

### 1. If golang is installed in your environment
```
git clone https://github.com/boanlab/rb.git
cd rb/src
go build
```
### 2. Homebrew
Currently investigating an issue where the installation occasionally fails to complete.

```
brew tap boanlab/rb
brew install rb
```
After installation is complete, you can run the rb command to execute it. 
For example, you can run it like this:
```
rb -url=https://example.com -w=10 -r=1000
```

### 3. Otherwise, use the binary file in the release

#### 3.1 Linux
```
wget https://github.com/boanlab/rb/releases/download/v0.0.2/rb_linux_amd64
mv rb_linux_amd64 rb
chmod +x rb

./rb -url=http://example.com -w=10 -r=100
```


## Usage
```
$ ./rb -url=http://127.0.0.1:8080  -type=f -w=10 -r=1000 -t=60s  ## If you have installed it via homebrew, the command to execute it would be just "rb" instead of "./rb".

Welcome to rb

Running Benchmark with type=f url=http://127.0.0.1:8080, workers=10, total requests =1000, timeout=1m0stotal 


Sent 1000 requests in 1m0s
Average request time: 15.8309 ms
Request time standard deviation: 4.9787 ms


Response time histogram:
9.239 - 11.849 ms [149] ▄▄▄▄▄▄▄
11.849 - 14.460 ms [338] ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄
14.460 - 17.070 ms [196] ▄▄▄▄▄▄▄▄▄▄
17.070 - 19.680 ms [171] ▄▄▄▄▄▄▄▄▄
19.680 - 22.290 ms [78] ▄▄▄▄
22.290 - 24.901 ms [32] ▄▄
24.901 - 27.511 ms [11] ▄
27.511 - 30.121 ms [15] ▄
30.121 - 32.731 ms [1] 
32.731 - 35.342 ms [1] 
35.342 - 37.952 ms [1] 
37.952 - 40.562 ms [1] 
40.562 - 43.172 ms [1] 
43.172 - 48.396 ms [4] 
48.396 - 48.406 ms [1] 


Response Time Percentiles:
 10% in 11.456 ms
 25% in 12.709 ms
 50% in 14.592 ms
 75% in 17.980 ms
 90% in 20.993 ms
 95% in 22.929 ms
 99% in 28.925 ms


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
