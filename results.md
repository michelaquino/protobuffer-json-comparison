# Protocol Buffer para serialização de dados
A ideia é testar o uso de Protocol Buffer para serialização de dados e  comparar com JSON: 


## Performance
### Serialização/Deserialização
- Serialização
```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz

BenchmarkSerializeJSON-8                           42920             28149 ns/op
BenchmarkSerializeProtocolBuffer-8                 44228             27277 ns/op
BenchmarkSerializeProtocolBufferAsJSON-8           30870             40263 ns/op
```

- Deserialização
```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz

BenchmarkDeserializeJSON-8                        110302             10970 ns/op
BenchmarkDeserializeProtocolBuffer-8              562112              2198 ns/op
BenchmarkDeserializeProtocolBufferAsJSON-8         69072             18385 ns/op
```

### Requests
#### POST

##### ProtocolBuffer
```
Running 10s test @ http://localhost:8888/proto
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    36.22ms   24.47ms 249.28ms   71.60%
    Req/Sec    29.12     13.40   161.00     76.51%
  29326 requests in 10.10s, 2.10MB read
Requests/sec:   2903.42
Transfer/sec:    212.65KB
```

##### JSON
```
Running 10s test @ http://localhost:8888/json
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    79.52ms  169.71ms   1.09s    93.17%
    Req/Sec    29.16     13.22   110.00     76.06%
  26869 requests in 10.09s, 1.92MB read
Requests/sec:   2661.85
Transfer/sec:    194.96KB
```

##### ProtocolBuffer como JSON
```
Running 10s test @ http://localhost:8888/protojson
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    40.33ms   27.15ms 241.85ms   70.52%
    Req/Sec    26.16     12.45   100.00     63.87%
  26330 requests in 10.10s, 1.88MB read
Requests/sec:   2607.63
Transfer/sec:    190.99KB
```

#### GET

##### ProtocolBuffer
```
Running 10s test @ http://localhost:8888/proto
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    25.33ms   14.87ms 143.20ms   72.99%
    Req/Sec    40.86     14.23   171.00     74.51%
  41078 requests in 10.10s, 13.59MB read
Requests/sec:   4067.33
Transfer/sec:      1.35MB
```

##### JSON
```
Running 10s test @ http://localhost:8888/json
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    26.13ms   15.08ms 135.74ms   71.77%
    Req/Sec    39.44     13.27   171.00     77.17%
  39667 requests in 10.10s, 22.66MB read
Requests/sec:   3927.37
Transfer/sec:      2.24MB
```

##### ProtocolBuffer como JSON
```
Running 10s test @ http://localhost:8888/protojson
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    27.08ms   15.91ms 156.36ms   73.76%
    Req/Sec    38.24     13.48   141.00     75.38%
  38499 requests in 10.09s, 22.65MB read
Requests/sec:   3813.99
Transfer/sec:      2.24MB
```

## Tamanho do dado serializado no Redis
### ProtocolBuffer
```
MEMORY USAGE "address-PROTO"
(integer) 312
```

```
GET "address-PROTO"
"\np\n\bname 101\x12$2ed0bb14-6710-42a8-931f-69d7ae3d0a8e\x1a\temail 384\"\n\n\b15-99-18\"\x0b\n\a11-7-70\x10\x01\"\x0c\n\b18-64-85\x10\x02*\x0c\b\xfc\xf0\xc9\x85\x06\x10\xc4\xa7\xa8\x9e\x01\nq\n\bname 660\x12$25a4b343-98da-42b8-85fa-be3e6ea0e0d2\x1a\temail 410\"\n\n\b20-50-93\"\x0c\n\b88-56-96\x10\x01\"\x0c\n\b50-78-43\x10\x02*\x0c\b\xfc\xf0\xc9\x85\x06\x10\xb8\xc7\xaa\x9e\x01"
```

### JSON
```
MEMORY USAGE "address-JSON"
(integer) 568
```

```
GET "address-JSON"
"{\"people\":[{\"name\":\"name 992\",\"id\":\"59ed3f8a-ef20-48c6-a392-52663360658c\",\"email\":\"email 122\",\"phones\":[{\"number\":\"19-31-88\"},{\"number\":\"80-46-55\",\"type\":1},{\"number\":\"33-58-89\",\"type\":2}],\"last_updated\":{\"seconds\":1622308984,\"nanos\":121239469}},{\"name\":\"name 553\",\"id\":\"0287e61b-5f63-45a9-8dff-9ebae231f31b\",\"email\":\"email 783\",\"phones\":[{\"number\":\"82-23-38\"},{\"number\":\"91-36-36\",\"type\":1},{\"number\":\"62-58-71\",\"type\":2}],\"last_updated\":{\"seconds\":1622308984,\"nanos\":121285583}}]}"
```

### Protocol Buffer como JSON
```
MEMORY USAGE "address-PROTO-JSON"
(integer) 576
```

```
GET "address-PROTO-JSON"
"{\"people\":[{\"name\":\"name 284\",\"id\":\"6bd9b8a1-bf1b-4999-8bb2-11839ed5cc06\",\"email\":\"email 590\",\"phones\":[{\"number\":\"98-63-12\"},{\"number\":\"49-36-95\",\"type\":\"HOME\"},{\"number\":\"24-81-62\",\"type\":\"WORK\"}],\"lastUpdated\":\"2021-05-29T17:23:11.963003464Z\"},{\"name\":\"name 666\",\"id\":\"629fad5c-ce7b-4265-9dcc-be99aac9bf63\",\"email\":\"email 545\",\"phones\":[{\"number\":\"50-92-79\"},{\"number\":\"18-63-76\",\"type\":\"HOME\"},{\"number\":\"31-33-11\",\"type\":\"WORK\"}],\"lastUpdated\":\"2021-05-29T17:23:11.963026415Z\"}]}"
```
