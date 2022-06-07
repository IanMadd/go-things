package main

import "fmt"

// terraform accomplishes nothing
func terraform(planets [8]string) [8]string {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
	return planets
}

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

	planets = terraform(planets)
	fmt.Println(planets)

}
