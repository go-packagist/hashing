# hashing

[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/hashing)](https://goreportcard.com/report/github.com/go-packagist/hashing)
[![tests](https://github.com/go-packagist/hashing/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/hashing/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/hashing)](https://pkg.go.dev/github.com/go-packagist/hashing)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/hashing
```

## Usage

```go
package main

import (
    "github.com/go-packagist/hashing"
)

func main() {
	m := hashing.NewManager(&hashing.Config{
		Driver: "bcrypt",
	})
	
	m.Driver().Make("password")
	m.Driver("bcrypt").Make("password")
	m.Driver("md5").Make("password")
	m.Driver("bcrypt").Check("password", "hashed password")
	m.Driver("md5").Check("password", "hashed password")
}
```

## License

MIT
