package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

var (
	p1  = Point{1, 2}
	p2  = Point{X: 1}
	p3  = Point{}
	pt1 = &Point{1, 2}
)

func main() {
	fmt.Println(Point{1, 2})

	p := Point{3, 4}
	p.X = 4 // 結構體字段用點號存取
	fmt.Println(p, p.Y)

	pt := &p
	pt.X = 5    // 使用隱式間接引用，直接寫 `pt.X`
	(*pt).Y = 6 // 通過 `(*pt).Y` 來訪問其字段 `Y`
	fmt.Println(p)

}
