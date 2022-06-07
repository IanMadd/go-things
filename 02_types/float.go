package main

import "fmt"

func main() {
	third := 1.0 / 3.0
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)
	fmt.Printf("%013.2f\n", third)

	numbers := 0015.1021
	fmt.Printf("%09.4f\n", numbers)
	shouldBeTrue := third*3 == 1
	fmt.Println(shouldBeTrue)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Printf("%05.2f\n", piggyBank)

	celsius := 21.0
	fmt.Println("Multiple then divide.")
	fmt.Print((celsius/5.0*9.0)+32, "° F\n")
	fmt.Print((9.0/5.0*celsius)+32, "° F\n")

	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Print(fahrenheit, "° F\n")

	otherBank := 0.0
	for i := 0; i < 11; i++ {
		otherBank += 0.1
		// fmt.Println(i, otherBank)
		fmt.Printf("%02d, %018.16f\n", i, otherBank)
	}
	fmt.Println(otherBank)
}
