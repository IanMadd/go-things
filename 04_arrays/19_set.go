package main

import (
	"fmt"
	"sort"
)

func is_true(m map[float64]bool, t float64) {
	if m[t] {
		fmt.Println("set member")
	} else {
		fmt.Println("not set member")
	}
}

func main() {
	var temperatures = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)

	for _, t := range temperatures {
		set[t] = true
	}

	is_true(set, -28.0)

	is_true(set, -32.0)

	fmt.Println(set)

	unique := make([]float64, 0, len(set))

	fmt.Println(unique)
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)

	fmt.Println(unique)
}
