package main

import "fmt"

const abeba = "ABEBA"

func main() {
	name := "Dmytro"
	fmt.Println("Hello " + name)

	const aboba = "Aboba"

	var str string
	fmt.Println(str)

	var (
		firstName = "Bob"
		lastName  = "Sendrik"
	)
	fmt.Println(firstName, lastName)

	var c = fmt.Sprintf("Hello %s", lastName) // string concatenation
	fmt.Println(c)

	str2 := test(32)
	fmt.Println(str2)

	a1, b1, c1 := test2()
	fmt.Println(a1, b1, c1)

	sum := 12
	sum2 := 0
	for i := 0; i < sum; i++ {
		sum2 += i
	}
	sum3 := 0

	for sum3 < 100 {
		sum3 += 1
	}

	val := 101
	if val < 100 {
		fmt.Println("Yep")
	} else if val > 100 {
		fmt.Print("NO")
	}

	i := 0
	switch i {
	case 1:
		fmt.Print(1)
	case 2:
		fmt.Print(2)
	default:
		fmt.Print(0)
	}
}

func test(age int) string {
	var resul = fmt.Sprintf("You are %d years old", age)
	return resul
}

type s struct {
	a string
	b string
	c string
}

func test2() (string, string, string) {
	a := "a"
	b := "b"
	c := "c"
	return a, b, c
}

// func test3() (a, b, c string) {
// 	a = "a"
// 	b = "b"
// 	c = "c"
// 	return
// }

func test4() {
	defer fmt.Print("first") // отложенное действие, выполняется перед выходом из метода, LIFO!!!
	fmt.Print("second")
}
