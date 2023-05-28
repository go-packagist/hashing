package main

import (
	"fmt"
	"github.com/go-packagist/hashing"
)

func main() {
	// Example1: use md5 hasher
	md5 := hashing.Md5Hasher{} // OR md5 := hashing.NewMd5Hasher()
	md5Hashed, err := md5.Make("123456")
	if err != nil {
		panic(err)
	}
	fmt.Println(md5Hashed)                      // e10adc3949ba59abbe56e057f20f883e
	fmt.Println(md5.Check("123456", md5Hashed)) // true

	// Example2: use bcrypt hasher
	bcrypt := hashing.BcryptHasher{} // OR bcrypt := hashing.NewBcryptHasher()
	bcryptHashed, err := bcrypt.Make("123456")
	if err != nil {
		panic(err)
	}
	fmt.Println(bcryptHashed)
	fmt.Println(bcrypt.Check("123456", bcryptHashed)) // true

	// Example3: use manager
	manager := hashing.NewManager(&hashing.Config{
		Driver: "bcrypt", // OR Driver: "md5"
	})

	// use default driver
	fmt.Println(manager.Driver().MustMake("123456"))
	fmt.Println(manager.Driver().Check("123456", manager.Driver().MustMake("123456"))) // true

	// use specified driver
	fmt.Println(manager.Driver("md5").MustMake("123456"))
	fmt.Println(manager.Driver("md5").Check("123456", manager.Driver("md5").MustMake("123456"))) // true
	fmt.Println(manager.Driver("bcrypt").MustMake("123456"))
	fmt.Println(manager.Driver("bcrypt").Check("123456", manager.Driver("bcrypt").MustMake("123456"))) // true
}
