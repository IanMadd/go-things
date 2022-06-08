package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice),
		cap(slice), slice)
}

func newItem(items []string, newItems ...string) []string {

	for i := range newItems {
		items = append(items, newItems[i])
		dump("Slice", items)
	}

	return items

}

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println(dwarfs)
	dwarfs = newItem(dwarfs, "Orcus", "Ceres", "Salacia", "Quaoar", "Sedna", "2007 OR10", "2002 MS4")

	fmt.Println(dwarfs)
}
