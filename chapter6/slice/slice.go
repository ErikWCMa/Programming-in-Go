package main

import (
	"fmt"
	"strings"
)

func change(c, d []string) {
	d[0] = "ddd"
	fmt.Println(c, d)
}

func main() {
	// 切片是數组的引用
	names := [4]string{"john", "paul", "george", "ringo"}
	fmt.Println(names) // [john paul george ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [john paul] [paul george]
	change(a, b)
	b[0] = "xxx"
	fmt.Println(a, b, names) // [john xxx] [xxx george] [john xxx george ringo]

	func(a, b []string) {
		b[0] = "eee"
		fmt.Println(a, b)
	}(a, b)
	fmt.Println(a, b)
	// 切片預設起始是0，預設結束是數組長度
	i := []int{1, 2, 3, 4}
	fmt.Println(i[0:4]) // [1 2 3 4]
	fmt.Println(i[:4])  // [1 2 3 4]
	fmt.Println(i[0:])  // [1 2 3 4]
	fmt.Println(i[:])   // [1 2 3 4]

	// 長度與容量
	fmt.Printf("len=%d, cap=%d, %v\n", len(i[:4]), cap(i[:4]), i[:4])    // len=4, cap=4, [1 2 3 4]
	fmt.Printf("len=%d, cap=%d, %v\n", len(i[2:4]), cap(i[2:4]), i[2:4]) // len=2, cap=2, [3 4]

	// nil 切片
	var j []int
	fmt.Printf("len=%d, cap=%d, %v\n", len(j), cap(j), j) // len=0, cap=0, []

	// make 創建切片
	aM := make([]int, 5)
	bM := make([]int, 0, 5)
	fmt.Printf("len=%d, cap=%d, %v\n", len(aM), cap(aM), aM) // len=5, cap=5, [0 0 0 0 0]
	fmt.Printf("len=%d, cap=%d, %v\n", len(bM), cap(bM), bM) // len=0, cap=5, []

	// 包含切片的切片
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// 向切片追加元素
	var sP []int
	fmt.Printf("len=%d, cap=%d, %v\n", len(sP), cap(sP), sP) // len=0, cap=5, []

	sP = append(sP, 1)
	sP = append(sP, 2)
	fmt.Printf("len=%d, cap=%d, %v\n", len(sP), cap(sP), sP) // len=0, cap=5, []

}
