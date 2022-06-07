package main

import "fmt"

func main() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println(dwarfs)

	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	fmt.Println(len(planets))
	fmt.Println(planets)

	for planet := 0; planet < len(planets); planet++ {
		fmt.Println(planets[planet])
	}

	for i := len(planets); i > 0; i-- {
		fmt.Println(planets[i-1])
	}

	for i, dwarf := range dwarfs {
		fmt.Println(dwarf, i)
	}

	for _, dwarf := range dwarfs {
		fmt.Println(dwarf)
	}
}
