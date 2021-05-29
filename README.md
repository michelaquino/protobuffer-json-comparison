# protobuffer-json-comparison
Comparison betewwen Protocol Buffer and JSON

## Dependencies
- [Go 1.16+](https://golang.org/)
- [GNU Make](https://www.gnu.org/software/make/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker compose](https://docs.docker.com/compose/)
- [WRK](https://github.com/wg/wrk)
- [Protocol Buffer compiler](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation)
- [Go protocol buffers plugin](https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go)

## Compile proto files
Following the instructions [here](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers) to install the Protocol Buffer compiler and the  Go protocol buffers plugin

To build the `.proto` files:
```
make compile-proto-files
```

## Run
### With docker
```
make run-as-docker
```

### Without docker
Need redis runnning locally

Download dependencies:
```
make setup
```

Running:
```
make run
```

## Endpoints
- /proto
- /json
- /protojson

### GET
```
curl http://localhost:8888/<endpoint>
```

### POST
```
curl -X POST http://localhost:8888/<endpoint>
```

### Example
```
curl http://localhost:8888/proto
curl -X POST http://localhost:8888/proto
```

## Benchmark Tests
```
make benchmark-test
```

## Performance tests with WRK
### POST method
```
make wrk-post endpoint=<endpoint>
```
### GET method
```
make wrk-get endpoint=<endpoint>
```

### Example
```
make wrk-post endpoint=proto
make wrk-get endpoint=proto
```

## Referencies
- https://developers.google.com/protocol-buffers/docs/gotutorial
- https://github.com/protocolbuffers/protobuf#protocol-compiler-installation
- https://seb-nyberg.medium.com/customizing-protobuf-json-serialization-in-golang-6c58b5890356
