package main

import "fmt"

func main() {
	var planets [8]string

	fmt.Println(len(planets) == 8)
	fmt.Println(planets[0] == "")
	planets[0] = "Mercury"

	fmt.Println(planets)

	planets[1] = "Venus"
	planets[2] = "Earth"

	fmt.Println(planets)

	earth := planets[2]
	fmt.Println(earth)

	fmt.Println(len(planets))
	fmt.Println(planets[3] == "")

	// planets[11] = "Pluto" produces runtime error
}
