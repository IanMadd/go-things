package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func calibration() sensor {
	return rand.Intn(50)
}

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		sensorOffset := s() + offset
		fmt.Println(sensorOffset)
		return sensorOffset
	}
}

func main() {

	count := 10
	for i := 0; i < count; i++ {
		sensor := calibrate(realSensor, calibration())
		fmt.Println(sensor())
	}

}
