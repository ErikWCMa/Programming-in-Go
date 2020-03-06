package main

import "fmt"

var ai = uint(10)

func main() {
	/* Closure */
	// x := 0
	// y := 0
	// incr := func() int {
	// 	fmt.Println("before++ x= ", x)
	// 	x++
	// 	fmt.Println("after++ x= ", x)
	// 	fmt.Println("y=", y)
	// 	return x
	// }
	// fmt.Println(incr())
	// fmt.Println("x= ", x)
	// fmt.Println(incr())
	// fmt.Println("x= ", x)
	//
	// nextEven := makeEvenGenerator()
	// fmt.Println(nextEven()) // 0
	// fmt.Println(nextEven()) // 2
	// fmt.Println(nextEven()) // 4
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(nextEven())
	// 	println("i in for= ", i)
	// 	anotherMakeEvenGenerator()
	// }

	/* Recursion */
	// fmt.Println(factorial(5))

	/* defer */
	// defer second()
	// first()

	/* Panic& Recover */
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

// Closure
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		println("i bef= ", i)
		ret = i
		i += 2
		println("i aft= ", i)
		return
	}
}

// Closure
func anotherMakeEvenGenerator() uint {
	println("ai bef= ", ai)
	ai += 2
	println("ai aft= ", ai)
	return ai
}

// Recursion
func factorial(x uint) uint {
	if x == 7 {
		fmt.Println(x)
		return x
	}
	fmt.Println(x)
	return x * factorial(x+1)
}

// defer
func first() {
	fmt.Println("1st")
}

// defer
func second() {
	fmt.Println("2nd")
}
