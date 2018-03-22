# Zapz: [Uber's Zap](https://github.com/uber-go/zap) --> [logzio](https://github.com/logzio)

Zapz creates a zap.Logger that sends logs to logzio. 
 This creates a custom [WriterSync](https://github.com/uber-go/zap/blob/master/zapcore/write_syncer.go) that buffers data on disk and drains every 5 seconds.

 Zapz uses [logzio-go](https://github.com/dougEfresh/logzio-go) to transport logs via HTTP

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Go Report][report-img]][report]

## Installation 
```shell
$ go get -u github.com/dougEfresh/zapz
```

## Quick Start
 
 ```go
package main

import (
  "os"

  "github.com/dougEfresh/zapz"
    )

func main() {
  l, err := zapz.New(os.Args[1]) //logzio token required
  if err != nil {
    panic(err)
  }

  l.Info("tester")
  // Logs are buffered on disk, this will flush it
  if l.Sync() != nil {
      panic("oops")
  }
}
```




## Getting Started

### Get Logzio token
1. Go to Logzio website
2. Sign in with your Logzio account
3. Click the top menu gear icon (Account)
4. The Logzio token is given in the account page

## Usage
    
Set Debug level: `zapz.New(token, zapz.SetLevel(zapcore.DebugLevel))`
    
Set CustomEncoder config: `zapz.New(token, zapz.SetEncodeConfig(cfg))`
    
Set Custom log type:

A zap field, default is zap.String("type", "zap-logger") in the log to type

`zapz.New(token, zapz.SetType(logzType))`


## Examples
    



## Prerequisites

go 1.x

## Tests
    
```shell
$ go test -v

```


## Contributing
 All PRs are welcome

## Authors

* **Douglas Chimento**  - [dougEfresh][me]

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

  [logz java](https://github.com/logzio/logzio-java-sender)

### TODO 

[doc-img]: https://godoc.org/github.com/dougEfresh/zapz?status.svg
[doc]: https://godoc.org/github.com/dougEfresh/zapz
[ci-img]: https://travis-ci.org/dougEfresh/zapz.svg?branch=master
[ci]: https://travis-ci.org/dougEfresh/zapz
[cov-img]: https://codecov.io/gh/dougEfresh/zapz/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/dougEfresh/zapz
[glide.lock]: https://github.com/uber-go/zap/blob/master/glide.lock
[zap]: https://github.com/uber-go/zap
[me]: https://github.com/dougEfresh
[report-img]: https://goreportcard.com/badge/github.com/dougEfresh/zapz
[report]: https://goreportcard.com/report/github.com/dougEfresh/zapz