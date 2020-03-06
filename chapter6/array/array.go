package main

import (
	"fmt"
)

func main() {
	array1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array1=", array1)

	array2 := [2][2]float32{{1.1, 2.2}, {3.3, 4.4}}
	fmt.Println("array2=", array2)

	array3 := [3][3][3]int{{{1}, {4, 5, 6}, {7, 8, 9}}, {{11, 12, 13}, {14, 15}, {17, 18, 19}}, {{21, 22, 23}, {24, 25, 26}, {27, 28, 29}}}
	fmt.Println("array3=", array3)
	fmt.Println("array3[2]=", array3[2])
	fmt.Println("array3[1][2]=", array3[1][2])

	array4 := [2][3][2][3]int{{{{1111, 1112, 1113}, {1121, 1122, 1123}}, {{1211, 1212, 1213}, {1221, 1222, 1223}}, {{1311, 1312, 1313}, {1321, 1322, 1323}}}, {{{2111, 2112, 2113}, {2121, 2122, 2123}}, {{2211, 2212, 2213}, {2221}}, {{2311, 2312, 2313}, {2321, 2322, 2323}}}}
	fmt.Println("array4=", array4)
	fmt.Println("array4[1]=", array4[1])
	fmt.Println("array4[0][1]=", array4[0][1])
	fmt.Println("array4[1][0][1]=", array4[1][0][1])
	fmt.Println("array4[0][1][0][1]=", array4[0][1][0][2])
}
