# uuid
A nice uuid implementation.

## Features

- Short(18 bytes)
- Letter+number only(double-click to copy).
- Auto increment.
- High performance.
- Unique but don't rely on random algorithm(max qps 238328 per-process)


## Usages

```

// UUID returns unique string by timestamp / IP / PID / autoincrement ID
uuid.UUID() 

// RandID returns a string base on the timestamp / rand int64
uuid.RandID() 

// Rand returns a given-size string base on the timestamp / rand int64
uuid.Rand(size int) 
```

## Performance
```
$ go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/defool/uuid
BenchmarkMyUUID-12               4792280               242 ns/op
BenchmarkGoogleUUID-12           2506358               470 ns/op
```
