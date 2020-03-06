package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32}

func main() {
	for i , v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	s := make([]int, 5)
	for i := range s {
		s[i] = 1 << uint(i)
	}
	for _, v := range s{
		fmt.Printf("%d\n", v)
	}
}