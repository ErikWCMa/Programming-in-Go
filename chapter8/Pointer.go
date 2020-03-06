package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 5
	y := 6
	zero(x)
	fmt.Println(x) // x is still 5

	zeroPointer(&x)
	fmt.Println(x)  // x is 0
	fmt.Println(&x) // x's memory address

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1

	fmt.Println(&y)
	temp := y
	*&y = x
	*&x = temp
	fmt.Println(y)
	fmt.Println(x)

	fmt.Println(reflect.TypeOf(&x))
	fmt.Println(reflect.TypeOf(*&x))
}

func zero(x int) {
	x = 0
}

func zeroPointer(y *int) {
	*y = 0
	fmt.Println("y=", y)
}

// new
func one(xPtr *int) {
	*xPtr = 1
}
