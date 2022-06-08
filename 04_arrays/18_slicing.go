package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice),
		cap(slice), slice)
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestrial := planets[0:4:4]
	other := planets[0:4]
	worlds := append(terrestrial, "Ceres")

	dump("terrestrial", terrestrial)
	dump("other", other)
	dump("worlds", worlds)
	dump("planets", planets)

	other_worlds := append(other, "Ceres")
	dump("other_worlds", other_worlds)
	dump("planets", planets)
}
