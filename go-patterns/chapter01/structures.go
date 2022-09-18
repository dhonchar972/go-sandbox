package main

import (
	"errors"
	"fmt"
	"sort"
)

type Speaking interface {
	sayHi()
	introduce(string) int
}

type Person struct {
	name string
	age  int
	pet  string
}

func (*Person) sayHi() {
	fmt.Println("HI!")
}

func (*Person) introduce(name string) int {
	fmt.Printf("HI!, my name is %s", name)
	return 0
}

func intro(p *Person) Speaking {
	p.introduce(p.name)
	return p
}

func intro2(p Person) Speaking {
	p.introduce(p.name)
	return &p
}

func intro3(p Speaking) Speaking {
	p.sayHi()
	return p
}

func addTo(base int, vals ...int) []int { // ...int - optionals
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func callStructures() {
	p := Person{"Bob", 26, "Ponky"}
	fmt.Println(p) // {Bob 26 Ponky}

	p2 := Person{
		"Tom",
		17,
		"Kiki",
	}
	fmt.Printf("Name: %v,\nAge: %v,\nBest frend: %v\n", p2.name, p2.age, p2.pet)

	arr := []Person{
		p,
		p2,
	}
	for i, v := range arr {
		fmt.Printf("Id: %v,\nName: %v,\nAge: %v,\nPet: %v\n", i, v.name, v.age, v.pet)
	}

	fmt.Println(p2.name)

	intro(&p2)

	var antonymPerson struct {
		name string
		age  int
		pet  string
	}
	antonymPerson.name = "bob"
	antonymPerson.age = 50
	antonymPerson.pet = "dog"
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet)
}

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

var opMap = map[string]func(int, int) int{ // function is an object!!!
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func functionAsParameter() {
	people := []Person{
		{"Patterson", 37, "Pat"},
		{"Tracy", 23, "Bobbert"},
		{name: "Fred", pet: "Fredson", age: 18},
	}
	sort.Slice(people, func(i int, j int) bool {
		return people[i].name < people[j].name
	})
	fmt.Println(people)
}

// defer - отложенное выполнение

// interfaces, maps and slices are reference types!!!!!!!!!!!!!!!!!!!!!!!!!!
