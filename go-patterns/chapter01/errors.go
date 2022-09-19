package main

import (
	"errors"
	"fmt"
	"os"
)

//hello, I am new godoc comment!!!
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0") // or // fmt.Errorf()
	}
	return numerator / denominator, numerator % denominator, nil
}

//callErrors - calls the execution of functions from errors.go, package main
func callErrors() {
	numerator := 20
	denominator := 3
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder, mod)
}
