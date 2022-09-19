package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Command string
}

func callTypesAndMethods() {
	m := Manager{Employee{"Bob", "12"}, "SomeCommand"}
	fmt.Println("Name:", m.Name)
}
