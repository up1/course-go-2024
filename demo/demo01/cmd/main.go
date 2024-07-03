package main

import (
	"fmt"
	"hello"
)

func main() {
	res := hello.SayHi()
	println(res)

	u := hello.NewUser("1", "John")
	fmt.Println(u)
}
