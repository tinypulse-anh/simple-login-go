# Simple session API in Golang

## Instructions

- `go get github.com/tinypulse-anh/simple-login-go`
- `cd ~/go/github.com/tinypulse-anh/simple-login-go`
- `go get`
- `cp .env.sample .env` and change `DATABASE_URL` to point to TINYpulse dev db.
- `go run main.go` _=> Listen to :3000 with 100 threads_

## Benchmarks

- Request: `GET http://localhost:3000/api/v1/sessions/new?username=khuyennguyen@test.com`
_(I have configured SSO login for that user to make the request more heavy.)_

### Concurrency = 1

```
> bombardier -c 1 -n 10000 http://localhost:3000/api/v1/sessions/new?username=khuyennguyen@test.com
Bombarding http://localhost:3000/api/v1/sessions/new?username=khuyennguyen@test.com with 10000 request(s) using 1 connection(s)
 10000 / 10000 [===================================================================] 100.00% 1468/s 6s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec      1513.91     117.31    1655.01
  Latency      657.90us   150.49us     9.13ms
  HTTP codes:
    1xx - 0, 2xx - 10000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:   459.73KB/s
```

### Concurrency = 100

```
> bombardier -c 100 -n 10000 http://localhost:3000/api/v1/sessions/new?username=khuyennguyen@test.com
Bombarding http://localhost:3000/api/v1/sessions/new?username=khuyennguyen@test.com with 10000 request(s) using 100 connection(s)
 10000 / 10000 [===================================================================] 100.00% 2934/s 3s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec      2961.33     797.63    5695.83
  Latency       33.04ms    31.16ms   263.62ms
  HTTP codes:
    1xx - 0, 2xx - 10000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:     0.88MB/s
```
