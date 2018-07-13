# As
[![Build Status](https://travis-ci.org/lunemec/as.svg?branch=master)](https://travis-ci.org/lunemec/as) [![Go Report Card](https://goreportcard.com/badge/github.com/lunemec/as)](https://goreportcard.com/report/github.com/lunemec/as) [![Maintainability](https://api.codeclimate.com/v1/badges/d0b5da039ba6172a1b3b/maintainability)](https://codeclimate.com/github/lunemec/as/maintainability) 
[![codecov](https://codecov.io/gh/lunemec/as/branch/master/graph/badge.svg)](https://codecov.io/gh/lunemec/as)


`As` is a library to convert numeric types with overflow check in Go.


## Instalation

`go get -u github.com/lunemec/as`

## Example
```go
package main

import (
    "fmt"

    "github.com/lunemec/as"
)

func main() {
	for _, n := range []int{127, 128} {
		num, err := as.Int8(n)
		if err != nil {
			fmt.Printf("Input invalid: %d, err: %s\n", num, err)
		} else {
			fmt.Printf("Input valid: %d\n", num)
		}
	}
	// Output: Input valid: 127
	// Input invalid: -128, err: 128 (int) overflows int8
}
```

## Architecture support
`As` currently supports these architectures and correctly handles overflows for 64/32 bit numbers:
* 386
* amd64
* arm
* arm64

It may work for s390x or ppc64le, but I have no way to test that.
