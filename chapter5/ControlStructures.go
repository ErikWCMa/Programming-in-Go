package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Print(i)
			fmt.Println(" even")
			i = i + 1
		} else {
			fmt.Print(i)
			fmt.Println(" odd")
			i = i + 1
		}
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}
