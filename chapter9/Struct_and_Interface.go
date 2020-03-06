package main

import (
	"fmt"
	"math"
)

var c Circle

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea1(x, y, r float64) float64 {
	return math.Pi * r * r
}

func circleArea2(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func info(x Shape) {
	fmt.Println(x.area())
}

/* method */
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return math.Pi * c.r * 2
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
func (m *MultiShape) perimeter() float64 {
	var perimeter float64
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func (p *Person) Talk() {
	fmt.Println("Hi, My name is", p.Name)
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

type Shape interface {
	area() float64
	perimeter() float64
}

type MultiShape struct {
	shapes []Shape
	// shapes [2]Shape
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea1(cx, cy, cr))

	c1 := Circle{x: 0, y: 0, r: 5}

	c2 := Circle{0, 0, 5}

	fmt.Println(c1.x, c1.y, c1.r)
	fmt.Println(c2.x, c2.y, c2.r)

	fmt.Println(circleArea2(c2))

	/* method */
	fmt.Println(c1.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	a := new(Android)
	a.Name = "Pie"
	a.Person.Talk()
	a.Talk()

	fmt.Println(totalArea(&c1, &r))

	fmt.Println(totalPerimeter(&c1, &r))

	fmt.Println(c1, r)
	fmt.Println(&c1, &r)
	fmt.Println(c1.area())
	m := MultiShape{[]Shape{&c1, &c2, &r}}
	fmt.Println(m.area())
	info(&c1)

}
