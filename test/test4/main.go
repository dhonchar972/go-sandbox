package main

import "fmt"

func someFunc(callback func(int, int) int, s string) int {
	result := callback(5, 1)
	fmt.Println(s)
	return result
}

func closureFunc(init int) func(int) int { //closure - замыкание
	sum := init
	return func(x int) int {
		sum += x
		return sum
	}
}

func closureFunc2(init int) func(int) int { //closure - замыкание
	sum := init
	return func(x int) int {
		x += sum
		return x
	}
}

type Point struct {
	X int
	Y int
}

func immutableMovePoint(p Point, x, y int) Point {
	p.X += x
	p.Y += y
	return p
}

func mutableMovePoint(p *Point, x, y int) {
	p.X += x
	p.Y += y
}

func (p *Point) move(x, y int) {
	p.X += x
	p.Y += y
}

func main() {
	sumCallback := func(h1, h2 int) int {
		return h1 + h2
	}

	someFunc(sumCallback, "do work")
	g := someFunc(sumCallback, "do work")
	fmt.Println(g)

	closure := closureFunc(1)
	fmt.Println("closures:")
	fmt.Println(closure(1))
	fmt.Println(closure(1))
	fmt.Println(closure(1))
	fmt.Println(closure(1))
	fmt.Println(closure(5))

	closure2 := closureFunc2(1)
	fmt.Println("closures:")
	fmt.Println(closure2(1))
	fmt.Println(closure2(1))
	fmt.Println(closure2(1))
	fmt.Println(closure2(1))
	fmt.Println(closure2(5))

	p := Point{1, 2}
	fmt.Println(p)

	pp := immutableMovePoint(p, 2, 4) // immutable function
	fmt.Println(pp)

	mutableMovePoint(&pp, 7, 9) // mutable function, we use reference
	fmt.Println(pp)

	pp.move(7, 3) // structur method
	fmt.Println(pp)
}
