package main

import "fmt"

type kelvin float64
type celsius float64
type fahrenheit float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit(k*1.8 - 459.67)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

// func celsiusToKelvin(c celsius) kelvin {
// 	return kelvin(c + 273.15)
// }

func main() {
	var k kelvin = 294.0
	var c1 celsius
	var c2 celsius

	fmt.Println(c1)
	c2 = k.celsius()
	var f1 fahrenheit = k.fahrenheit()
	fmt.Println(c2)
	fmt.Println(f1)
	fmt.Println(f1.celsius())
}
