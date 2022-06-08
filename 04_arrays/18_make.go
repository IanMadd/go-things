package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice),
		cap(slice), slice)
}

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	dwarfs := make([]string, 0, 10)

	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	dump("Dwarfs", dwarfs)

	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)

	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets)
}
