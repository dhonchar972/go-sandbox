package main

import "fmt"

// type error interface { // has default realisation in standard library
// 	Error() string
// }

type AppError struct {
	Err    error
	Custom string
	Field  int
}

func (e *AppError) Error() string {
	err := fmt.Errorf("new error %s", e.Custom)
	return err.Error()
}

func main() {
	err := m()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func m() error { // custom error
	return &AppError{
		Err:    fmt.Errorf("custom error"),
		Custom: "value here",
		Field:  89,
	}
}
