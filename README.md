# go-jsoneq

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/go-jsoneq/jsoneq)
[![Build Status](https://cloud.drone.io/api/badges/suzuki-shunsuke/go-jsoneq/status.svg)](https://cloud.drone.io/suzuki-shunsuke/go-jsoneq)
[![codecov](https://codecov.io/gh/suzuki-shunsuke/go-jsoneq/branch/master/graph/badge.svg)](https://codecov.io/gh/suzuki-shunsuke/go-jsoneq)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/go-jsoneq)](https://goreportcard.com/report/github.com/suzuki-shunsuke/go-jsoneq)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/go-jsoneq.svg)](https://github.com/suzuki-shunsuke/go-jsoneq)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/go-jsoneq.svg)](https://github.com/suzuki-shunsuke/go-jsoneq/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/go-jsoneq/master/LICENSE)

Golang library to check if two values are equal as JSON.

```go
package jsoneq_test

import (
	"fmt"
	"log"

	"github.com/suzuki-shunsuke/go-jsoneq/jsoneq"
)

type (
	Foo struct {
		Foo string `json:"foo"`
	}
)

func checkResult(b bool, err error) {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}

func Example() {
	b, err := jsoneq.Equal(
		Foo{Foo: "bar"},
		map[string]string{"foo": "bar"},
	)
	checkResult(b, err)

	// when a type of value is []byte, it is treated as JSON string
	b, err = jsoneq.Equal(
		Foo{Foo: "bar"},
		[]byte(`{"foo": "bar"}`),
	)
	checkResult(b, err)

	// Output:
	// true
	// true
}
```

## License

[MIT](LICENSE)
