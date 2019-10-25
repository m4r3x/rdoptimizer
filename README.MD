# Purpose
To test and play around with Redis performance.

# Pre-Requirements:
1.  Have a running redis server, you can easy boot one using docker: `docker run -p 127.0.0.1:6379:6379 -d redis`
2.  (not necessary, proto objects are already built) Have Google's protobuf installed (necessary for compilation of .proto files) and compile message: `cd proto && protoc --go_out=. *.proto`

# Environment variables:
- REDIS_CONNECTION - it's full URL redis connection strings, it defaults to `127.0.0.1:6379`

# Usage
Binary accepts two parameters, number of requests per second and benchmark type.

i.e:
- `go run rdoptimizer.go 5000 1` will run `rawEventsBenchmark` (type 1) with 5000 requests per seconds.
- `go run rdoptimizer.go 100 6` will run `protoTextBenchmark` (type 6) with 100 requests per seconds.

All types are present in `benchmark.go` switch statement.

Sample app output looks like that:
```
Redis connection: localhost:6379
Redis connection: OK
Redis flush: OK
RPS: 5000
Raw text scenario, case study #2 - proto encoded data.

Redis Alloc = 11260 KB  Host Alloc = 18 MB      Cycle = 1
Redis Alloc = 23579 KB  Host Alloc = 24 MB      Cycle = 2
Redis Alloc = 32919 KB  Host Alloc = 14 MB      Cycle = 3
```