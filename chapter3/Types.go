package main

import "fmt"

var hello string = "Hello World"

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[9])
	fmt.Println(hello[2:4])
	fmt.Println("Hello " + "World")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
