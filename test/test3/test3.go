package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Point struct {
	X int
	Y int
}

func main() {
	arr := []string{"a", "b", "c", "d"}
	for i, j := range arr {
		fmt.Printf("Position: %d, Value: %s\n", i, j)
	}

	for _, j := range arr {
		fmt.Printf("Value: %s\n", j)
	}

	pointsMap := map[string]Point{}
	pointsMap2 := make(map[string]Point)
	pointsMap["a"] = Point{
		X: 1,
		Y: 2,
	}
	pointsMap2["b"] = Point{X: 2, Y: 3}
	pointsMap2["c"] = Point{4, 5}
	fmt.Printf("X: %d\n", pointsMap2["c"].X)

	value, ok := pointsMap["a"]
	if ok {
		fmt.Printf("key=%s exist in map\n", "a")
		fmt.Println(value)
	} else {
		fmt.Printf("key=%s doesn't exist in map\n", "a")
		fmt.Println(value)
	}

	p1 := Point{}
	mapstructure.Decode(pointsMap["a"], &p1)
	fmt.Println(p1)
}
