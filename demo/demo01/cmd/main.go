package main

import (
	"fmt"
	"hello"
	"unsafe"
)

func main() {

	type first struct {
		f float64
		i int32
		b bool
	}
	a := first{}
	fmt.Println(unsafe.Sizeof(a))

	res := hello.SayHi()
	println(res)

	u := hello.NewUser("1", "John")
	res2 := u.GetById()
	println(res2)
	println(u.Id)

	u2 := hello.NewUserV2("2", "Doe", 30)
	u2.Name = "Jane"
	u2.Id = "3"
	u2.Age = 25
	u2.GetById()
}
