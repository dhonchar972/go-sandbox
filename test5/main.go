package main

import "fmt"

type Some interface { // interface that must be implemented
	Full() string
}

type some struct { // structure, essentially class fields
	Val1, Val2 string
}

func NewSome(val1, val2 string) Some { // class constructor
	s := some{
		Val1: val1,
		Val2: val2,
	}
	return &s
}

func (s *some) Full() string { // class method
	return s.Val1 + s.Val2
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	var inter Some = &some{"4", "7"}
	fmt.Println(inter.Full())
	var a Some
	var b *some
	fmt.Printf("Value:%v, Type:%T\n", a, a)
	fmt.Printf("Value:%v, Type:%T\n", b, b)
	a = b
	fmt.Printf("Value:%v, Type:%T\n", a, a)

	var i interface{} // mutable value
	i = "dadadad"
	i = 4
	fmt.Println(i)
	k, ok := i.(float32)
	fmt.Println(k, ok) // 0 false, can't cast to type float

	var h interface{} = "ghhgjgjgj"
	t, ok := h.(string)
	fmt.Println(t, ok) // ghhgjgjgj true
	c, ok := h.(float64)
	fmt.Println(c, ok) // 0 false, can't cast to type float
}
