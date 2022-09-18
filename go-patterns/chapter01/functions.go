package main

import (
	"fmt"
	"math/rand"
	"time"
)

func function1() {
	x := 10
	if x > 5 {
		fmt.Println(x) // 10
		x := 5
		fmt.Println(x) // 5
	}
	fmt.Println(x) // 10
}

func callFunctions() {
	function1()

	fmt.Println(true)
	true := 10
	fmt.Println(true) // 10

	ifElse()
	forTypesOfFor()
}

func ifElse() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomNumber := r.Intn(100)
	fmt.Println("A random number: ", randomNumber)

	if n := rand.Intn(10); n == 0 { // initialization n
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}

func forTypesOfFor() {
	// full for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 1 // while
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	for { // unlimited while
		fmt.Println(i)
		//continue
		break
	}

	for {
		// do - while
		if i != 0 {
			break
		}
	}

	for i, u := range "hello" { // for-range(for-each)
		fmt.Println(" + ", i, "-", u)
	}

	it := [10]int{}
	k := 10
	for i := range it {
		it[i] = k
		k++
	}
	for _, i := range it {
		fmt.Println(i)
	}
	for i := range it {
		it[i] = k
		k++
	}
	for k, i := range it {
		fmt.Println(k, i)
	}

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals) // [2 4 6 8 10 12]
	for i := range evenVals {
		evenVals[i] *= 2
	}
	fmt.Println(evenVals) // [4 8 12 16 20 24]

	samples := []string{"hello", "apple_π!"}
outer: // mark
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer // break outer, goto outer
			}
		}
		fmt.Println()
	}

	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}
