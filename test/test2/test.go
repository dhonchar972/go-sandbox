package main

import (
	"fmt"
)

func main() {
	a := "Hello"
	b := &a
	fmt.Println(b)  // 0xc000044250
	fmt.Println(*b) // Hello

	*b = "World"
	fmt.Println(b)  // 0xc000044250
	fmt.Println(*b) // World
	fmt.Println(a)  // World

	g := structs(12, 7)
	ref := &g
	fmt.Printf("X:%d, Y:%d\n", ref.X, ref.Y) // X:12, Y:7
	fmt.Println(ref)                         // &{12 7} struct reference
	fmt.Println(&ref)                        // 0xc00000a030 memory address
	fmt.Println(*ref)                        // {12 7} struct

	ref.method()

	arr := [...]string{"fsfsfsfs", "wfwfwfwf"} // array
	fmt.Println(arr)
	slice := []string{ // slice
		"fsfsfsfs",
		"wfwfwfwf",
	}
	slice = append(slice, "new")
	fmt.Println(slice)

	createSlice := make([]int, 5) // slice, something like list in java
	fmt.Println(createSlice)

	// reference type, changes in a slice change the data in the original and in other slices
	var slice2 []string = slice[0:1]
	fmt.Println(slice2)
}

type Point struct {
	X int
	Y int
}

func structs(x int, y int) Point {
	p := Point{
		X: x,
		Y: y,
	}
	return p
}

func (p *Point) method() {
	fmt.Printf("X:%d, Y:%d\n", p.X, p.Y)
}
