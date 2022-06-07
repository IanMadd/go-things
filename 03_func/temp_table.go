package main

import (
	"fmt"
	"strings"
)

type temp float64
type convert func(temp) temp

var bars = strings.Repeat("=", 23)

func convertCToF(c temp) temp {
	return c*1.8 + 32
}

func convertFToC(f temp) temp {
	return (f - 32) / 1.8
}

func printTempConversionTable(c convert, minTemp temp, maxTemp temp, tempArray [2]string) {
	fmt.Println(bars)
	fmt.Printf("| °%s       | °%s       |\n", tempArray[0], tempArray[1])
	fmt.Println(bars)
	for temp := minTemp; temp <= maxTemp; temp += 5 {
		convertedTemp := c(temp)
		fmt.Printf("| %-8.1f | %-8.1f |\n", temp, convertedTemp)
	}
	fmt.Println(bars)
}

func main() {
	var tempArray = [2]string{"C", "F"}
	printTempConversionTable(convertCToF, temp(-40.0), temp(100.0), tempArray)
	tempArray = [2]string{"F", "C"}
	printTempConversionTable(convertFToC, temp(-40.0), temp(100.0), tempArray)

}
