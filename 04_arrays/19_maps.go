package main

import "fmt"

func print_moon(temp map[string]int) {
	if moon, ok := temp["Moon"]; ok {
		fmt.Printf("On average the moon is %v° C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
}

func main() {

	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	something := map[float64]int{}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %v° C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	print_moon(temperature)
	moon := temperature["Moon"]
	fmt.Println(moon)
	temperature["Moon"] = moon
	print_moon(temperature)

}
