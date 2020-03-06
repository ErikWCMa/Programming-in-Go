package main

import "fmt"

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	var y string 
	y = "Hello World"
	fmt.Println(y)

	var z string
	z = "first"
	fmt.Println(z)
	z = "second"
	fmt.Println(z)

	z = "first"
	fmt.Println(z)
	z = z + "second"
	fmt.Println(z)

	fmt.Println(x == y)
	fmt.Println(x == z)

	w := 5
	fmt.Println(w)
	
	fmt.Print("Enter a number: ")
    var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
    
    fmt.Println(output)   
}