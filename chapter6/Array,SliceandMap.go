package main

import "fmt"

func main() {
    var x [5]int
    x[4] = 100
	fmt.Println(x)
	
	fmt.Println(x[4])

	var y [5]float64
    y[0] = 98
    y[1] = 93
    y[2] = 77
    y[3] = 82
    y[4] = 83
    var total float64 = 0
    for i := 0; i < 5; i++ {
        total += y[i]
    }
    fmt.Println(total / 5)

	var total2 float64 = 0
    for i := 0; i < len(y); i++ {
        total2 += y[i]
    }
	fmt.Println(total2 / float64(len(y)))
	
	c := [6]string{"a","b","c","d","e","f"}
	fmt.Println(c[2:5])


	f := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	fmt.Println(findMin(f))
	fmt.Println(f)
	
	xs := []int{1,2,3}
	fmt.Println(add(xs...))
	
	add := func(x, y int) int {
        return x + y
    }
    fmt.Println(add(1,1))  
}


func findMin(a []int)(int, int) {
	min := a[0]
	max := a[0]
	for i:=0 ;i < len(a);i++ {
		if min > a[i]{
			min = a[i]
		}
		if max < a[i]{
			max = a[i]
		}
	}
	return min,max
}

func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}