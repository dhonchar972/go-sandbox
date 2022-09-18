package main

import "fmt"

func callMaps() {
	//var nilMap map[string]int
	//nilMap["boy"] = 4
	//fmt.Println(nilMap["boy"]) // PANIC!!!!

	totalWins := map[string]int{}
	totalWins["boy"] = 4
	fmt.Println(totalWins["boy"]) // 4
	totalWins["boy"]++
	fmt.Println(totalWins["boy"]) // 5
	totalWins["k"]--
	fmt.Println(totalWins["k"]) // -1

	map1 := map[string]string{
		"one":   "hehe",
		"two":   "hihi",
		"three": "aaaa",
	}
	fmt.Println(map1) // map[one:hehe three:aaaa two:hihi]
	map1["four"] = "kek"
	fmt.Println(map1["four"]) // kek

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)               // map[Kittens:[Waldo Raul Ze] Lions:[Sarah Peter Billie] Orcas:[Fred Ralph Bijou]]
	fmt.Println(teams["Kittens"][0]) // Waldo

	fixmap := make(map[string]float64, 10) // Preallocate room for 10 entries
	fmt.Println(fixmap)                    // map[]

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"] // 5 true
	fmt.Println(v, ok)
	v, ok = m["world"] // 0 true
	fmt.Println(v, ok)
	v, ok = m["goodbye"] // 0 false
	fmt.Println(v, ok)

	delete(m, "hello")
	fmt.Println(m["hello"]) // 0

	// something like set(emulation)
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), " ", len(intSet)) // 11 8
	fmt.Println(intSet[5])                   // true
	fmt.Println(intSet[500])                 // false
	if intSet[7] {
		fmt.Println("7 is in the set")
	}

	// sad alternative
	intSet2 := map[int]struct{}{}
	vals2 := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals2 {
		intSet2[v] = struct{}{}
	}
	if _, ok := intSet2[5]; ok {
		fmt.Println("5 is in the set")
	}
}
