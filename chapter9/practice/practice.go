package main

import "fmt"
import "math"

type Circle struct {
	x, y, r float64
}

func circleArea1(c Circle) float64 {
	if c.r >9 {
		c.r=5
	}
    return math.Pi * c.r *c.r
}

func circleArea2(c *Circle) float64 {
	if c.r >9 {
		c.r=5
	}
    return math.Pi * c.r *c.r
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)
	c.r = 10
	fmt.Println(circleArea1(c))	
	fmt.Println(c.x, c.y, c.r)	
	fmt.Println(circleArea2(&c))
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(&c.x, &c.y, &c.r)	
	fmt.Println(&c)	

}

// func distance(x1, y1, x2, y2 float64) float64 {
// 	a := x2 – x1
// 	b := y2 – y1
// 	return math.Sqrt(a*a + b*b)
// }

// func rectangleArea(x1, y1, x2, y2 float64) float64 {
// 	l := distance(x1, y1, x1, y2)
// 	w := distance(x1, y1, x2, y1)
// 	return l * w  
// }
 
// func circleArea(x, y, r float64) float64 {
// 	return math.Pi * r*r
// }
