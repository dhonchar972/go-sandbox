package main

import "fmt"

type pet struct {
	name string
	age  int
}

type human struct {
	name string
	age  int
	pet  pet
}

type reachHuman struct {
	name string
	age  int
	pets []pet
}

func callEmbeddedStructure() {
	h := human{
		name: "Bob",
		age:  17,
		pet: pet{
			name: "Pim",
			age:  2,
		},
	}

	h2 := human{"Tom", 13, pet{"Lim", 6}}
	fmt.Println(h, h2) // {Bob 17 {Pim 2}} {Tom 13 {Lim 6}}

	rh := reachHuman{
		"Daniel",
		22,
		[]pet{
			pet{
				"Bonny",
				4,
			},
			pet{
				"Kobo",
				5,
			},
			pet{
				"Rio",
				7,
			},
		},
	}
	for i, v := range rh.pets {
		fmt.Println(i, v)
		//0 {Bonny 4}
		//1 {Kobo 5}
		//2 {Rio 7}
	}
}
