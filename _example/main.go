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
