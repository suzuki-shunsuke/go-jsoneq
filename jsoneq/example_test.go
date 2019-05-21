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

	// use jsoneq.Byte to check if a value is equal to JSON string
	b, err = jsoneq.Equal(
		Foo{Foo: "bar"},
		jsoneq.Byte([]byte(`{"foo": "bar"}`)),
	)
	checkResult(b, err)

	// Output:
	// true
	// true
}
