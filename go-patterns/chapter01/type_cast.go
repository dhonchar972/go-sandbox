package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func runTypesCast() {
	v := "1984"
	// string, numeral system(binary, decimal, hexadecimal, ...), bit size(0 - int(int64), 32 - int32, 64 - int64)
	m, _ := strconv.ParseInt(v, 10, 0)
	// integer, numeral system
	l := strconv.FormatInt(m, 10)
	fmt.Printf("%d,%s\n", m, reflect.TypeOf(m)) // 1984,int64?????????WHY????????
	fmt.Printf("%s,%s\n", l, reflect.TypeOf(l)) // 1984,string

	var i int = 34
	s := strconv.Itoa(i)
	p, _ := strconv.Atoi(s)
	fmt.Printf("%d,%s\n", i, reflect.TypeOf(i)) // 34,int
	fmt.Printf("%s,%s\n", s, reflect.TypeOf(s)) // 34,string
	fmt.Printf("%d,%s\n", p, reflect.TypeOf(p)) // 34,int
}
