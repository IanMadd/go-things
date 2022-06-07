package main

import (
	"fmt"
)

type celsius float64

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k celsius) celsius {
	k -= 273.15
	return k
}

func main() {
	var temperature celsius = 20.0
	fmt.Println(temperature)
	var temp celsius = 20.0
	tempCelsius := kelvinToCelsius(temp)
	fmt.Println(tempCelsius)
}
