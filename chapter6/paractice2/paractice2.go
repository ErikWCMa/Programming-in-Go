package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func change(c, d []string) {
	d[0] = "ddd"
	fmt.Println(c, d)
}

func arr(x [2]int) {
	fmt.Printf("pass Array : %p , %v\n", &x, x)
}

func main() {
	names := [4]string{"john", "paul", "george", "ringo"}
	fmt.Println(names) // [john paul george ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [john paul] [paul george]
	change(a, b)
	b[0] = "xxx"             // 更改切片的元素会修改其底层数组中对应的元素，并且与它共享底层数组的切片都会观测到这些修改
	fmt.Println(a, b, names) // [john xxx] [xxx george] [john xxx george ringo]
}
