# Purpose
To test and play around with Redis performance.

# Pre-Requirements:
1.  Have a running redis server, you can easy boot one using docker: `docker run -p 127.0.0.1:6379:6379 -d redis`
2.  Have Google's protobuf installed (necessary for compilation of .proto files) and compile message: `cd proto && protoc --go_out=. *.proto`

# Environment variables:
- REDIS_CONNECTION - it's full URL redis connection strings