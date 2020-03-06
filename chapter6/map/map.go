package main

import (
	"fmt"
	"reflect"
)

type vertex struct {
	Lit, Log float64
}

var m map[string]vertex

var m1 = map[string]vertex{
	"Bell Labs": vertex{40.123, -74.123}, "Google": vertex{37.123, -122.123},
}

var m2 = map[string]vertex{
	"Bell Labs": {40.123, -74.123},
	"Google":    {37.123, -122},
}

func main() {
	m = make(map[string]vertex)
	m["Bell Labs"] = vertex{40.123, -74.123}
	fmt.Println(m["Bell Labs"], m2)

	user := make(map[string]int)
	user["Age"] = 42
	fmt.Println(user["Age"])
	reflect.TypeOf(user)
	v, ok := user["Age"]
	fmt.Println("The Age:", v, "Presen?", ok)
}
