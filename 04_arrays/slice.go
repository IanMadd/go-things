package main

import (
	"fmt"
	"reflect"
	"sort"
)

type StringSlice []string

// func (p StringSlice) Sort()

func main() {
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

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println(terrestrial, gasGiants, iceGiants)
	fmt.Println(gasGiants[0])
	fmt.Println(reflect.TypeOf(terrestrial))
	fmt.Println(fmt.Sprintf("%T", terrestrial))
	fmt.Println(reflect.TypeOf(planets))

	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]
	fmt.Println(giants, gas, ice)

	neptune := "Neptune"
	tune := neptune[3:]

	fmt.Println(tune)

	question := "¿Cómo estás?"
	fmt.Println(question[:6]) //indices indicate the number of bytes, not runes

	colonized := terrestrial[2:]
	fmt.Println(colonized)

	dwarfs := []string{
		"Ceres",
		"Pluto",
		"Haumea",
		"Makemake",
		"Eris",
	}
	fmt.Println(dwarfs)
	fmt.Printf("%T\n", dwarfs)

	animals := []string{
		"goats",
		"chickens",
		"pigs",
		"horses",
		"cats",
		"dogs",
	}

	sort.Strings(animals)
	fmt.Println(animals)

}
