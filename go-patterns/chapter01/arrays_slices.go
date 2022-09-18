package main

import (
	"fmt"
	"reflect"
)

/* we use "verbs" in this function, examples:
%t - boolean
%d - integer
%b - decimal
%g - float
%s - string
%p - pointer(address)
%v - value(pointer value, data) !!!!!!!
*/
func callArrays() {
	var x1 [3]int        // array with length 3, -> 0,0,0
	fmt.Println(len(x1)) // 3
	fmt.Println(cap(x1)) // 3
	fmt.Println(x1)      // [0 0 0]

	var x2 = [3]int{10, 20, 30} // array with length 3, -> 10, 20, 30
	fmt.Println(x2)             // [10 20 30]

	var x3 = [12]int{1, 5: 4, 6, 10: 100, 15} // array with gaps (empty cells) -> [1,0,0,0,0,4,6,0,0,0,100,15]
	fmt.Println(x3)                           // [1 0 0 0 0 4 6 0 0 0 100 15]

	var x4 = [...]int{10, 20, 30} // array with length autodetect(3), -> 10, 20, 30
	fmt.Println(len(x4))          // 3
	fmt.Println(x4)               // [10 20 30]

	var x5 [2][3]int                                          // multidimensional array(2x3)
	fmt.Printf("multidimensional array lanth: %d\n", len(x5)) // 2
	fmt.Printf("multidimensional array cap: %d\n", cap(x5))   // 2
	fmt.Printf("multidimensional array: %v\n", x5)            // [[0 0 0] [0 0 0]]

	a1 := [3]int{}
	a2 := [3]int{}
	b := a1 == a2
	fmt.Printf("a1 equals to a2: %t\n", b) // a1 equals to a2: true

	//a3 := [4]int{}
	//b2 := a1 == a3 // mismatched types!!!!

	a4 := [3]int{1, 2, 3}
	a5 := [3]int{0, 1, 2}
	fmt.Printf("a4 equals to a5: %t\n", a4 == a5) // false

	a6 := [3]int{1, 2, 3}
	a7 := [3]int{1, 2, 3}
	fmt.Printf("a6 equals to a7: %t\n", a6 == a7) // true

	a8 := [3]int{3, 2, 1}
	a9 := [3]int{1, 2, 3}
	fmt.Printf("a8 equals to a9: %t\n", a8 == a9) // false
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////
func callSlices() {
	var es []int // nil
	fmt.Println(es)
	//es = append(es, 4) // panic: runtime error: index out of range [0] with length 0

	s := []int{10, 20, 30}
	fmt.Println(s)

	s2 := []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(s2)

	//var s3 [][]int // panic: runtime error: index out of range [0] with length 0
	// var s3 [0][0]int // we must use append but its multidimensional slice, ammmmmmmmmmmmmmmmmm
	var s3 [1][1]int
	fmt.Println(s3) // [[0]]
	s3[0][0] = 10
	fmt.Println(s3)       // [[10]]
	fmt.Println(s3[0][0]) // 10
	fmt.Println(len(s3))  // 1
	fmt.Println(cap(s3))  // 1

	es = append(es, 10, 12, 13)
	s2 = append(s2, s...)
	fmt.Println(s2)      // [1 0 0 0 0 4 6 0 0 0 100 15 10 20 30]
	fmt.Println(len(s2)) // 15
	fmt.Println(cap(s2)) // 24
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////
func callSlicesWithMake() {
	m := make([]int, 5)      // резервирование места для среза, длина 5(0-4), инициализирует нулями!!!
	m = append(m, 10)        // !!!!!ВСЕГДА УВЕЛИЧИВАЕТ ДЛИНУ СРЕЗА, НЕ ДЕЛАЙ ТАК!!!!!!
	m2 := make([]int, 0, 10) // (тип, длина, емкость)
	m2 = append(m2, 10)      // так нормально, длина 0, емкость не увеличится!!!

	var data []int // -> nil
	fmt.Println(data)
	fmt.Println(data == nil) // true
	fmt.Println(reflect.TypeOf(data))

	data2 := []int{} // != nil, срез нулевой длины
	fmt.Println(data2)
	fmt.Println(data2 == nil) // false

	ss1 := m2[:2]                                                      // [10 0]
	ss2 := m2[1:]                                                      // []
	ss3 := m2[1:3]                                                     // [0 0]
	ss4 := m2[:]                                                       // [10]
	fmt.Printf("1 - %p\n2 - %p\n3 - %p\n4 - %p\n", ss1, ss2, ss3, ss4) // pointers
	fmt.Printf("1 - %v\n2 - %v\n3 - %v\n4 - %v\n", ss1, ss2, ss3, ss4) // values

	x := []int{1, 2, 3, 4}
	y := make([]int, 3)
	copy(y, x)
	for i, e := range y {
		fmt.Printf("Index: %v, Value: %v\n", i, e)
	}

	j := y[:]
	//z := y[0:4]
	//f := y[0:4:5]// срез с емкостью
	fmt.Printf("val: %v, cap: %v, len: %v\n", j, cap(j), len(j)) // val: [1], cap: 1, len: 1
	//fmt.Printf("val: %v, cap: %v, len: %v\n", z, cap(z), len(z)) // val: [1], cap: 1, len: 1
	//fmt.Printf("val: %v, cap: %v, len: %v\n", f, cap(f), len(f)) // val: [1], cap: 1, len: 1
	fmt.Println(len(y))
	fmt.Println(cap(y))

	y = append(y, 12, 9, 37, 22, 18, 4, 45, 78, 40, 34, 79, 34, 34, 44, 54, 87, 45, 8, 9, 54, 8)
	fmt.Printf("val: %v, cap: %v, len: %v\n", j, cap(j), len(j)) // val: [1], cap: 1, len: 1
	//fmt.Printf("val: %v, cap: %v, len: %v\n", z, cap(z), len(z)) // val: [1], cap: 1, len: 1
	//fmt.Printf("val: %v, cap: %v, len: %v\n", f, cap(f), len(f)) // val: [1], cap: 1, len: 1
	for i, e := range y {
		fmt.Printf("Index: %v, Value: %v\n", i, e)
	}
	fmt.Println(len(y))
	fmt.Println(cap(y))
}
