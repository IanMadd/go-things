package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// minYear := 1000
	// maxYear := 3050
	rand.Seed(time.Now().UTC().UnixNano())
	// year := rand.Intn(maxYear-minYear) + minYear
	year := 1900
	fmt.Println(year)
	if year%4 == 0 && year%100 == 0 && year%400 == 0 {
		fmt.Println("Leap year")
	}
	// if  {
	// 	fmt.Println("Divisible by 100")
	// }
	// if year%400 == 0 {
	// 	fmt.Println("Divisible by 100")
	// }
}
