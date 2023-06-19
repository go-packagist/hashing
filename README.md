# hashing

![Go](https://badgen.net/badge/Go/%3E=1.17/orange)
[![Go Version](https://badgen.net/github/release/go-packagist/hashing/stable)](https://github.com/go-packagist/hashing/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/hashing/v2)](https://pkg.go.dev/github.com/go-packagist/hashing/v2)
[![codecov](https://codecov.io/gh/go-packagist/hashing/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/hashing)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/hashing)](https://goreportcard.com/report/github.com/go-packagist/hashing)
[![tests](https://github.com/go-packagist/hashing/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/hashing/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/hashing/v2
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/go-packagist/hashing/v2"
)

func main() {
	// Example1: use md5 hasher
	md5Hashed, err := hashing.Md5.Make("123456")
	if err != nil {
		panic(err)
	}
	fmt.Println(md5Hashed)                              // e10adc3949ba59abbe56e057f20f883e
	fmt.Println(hashing.Md5.Check("123456", md5Hashed)) // true

	// Example2: use bcrypt hasher
	bcryptHashed, err := hashing.Bcrypt.Make("123456")
	if err != nil {
		panic(err)
	}
	fmt.Println(bcryptHashed)
	fmt.Println(hashing.Bcrypt.Check("123456", bcryptHashed)) // true

	// Example3: use sha1 hasher with cost
	sha1Hashed, err := hashing.Sha1.Make("123456")
	if err != nil {
		panic(err)
	}
	fmt.Println(sha1Hashed)
	fmt.Println(hashing.Sha1.Check("123456", sha1Hashed)) // true
}
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.