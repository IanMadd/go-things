package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// Spaceline        Days Trip type  Price
// ======================================
// Virgin Galactic    23 Round-trip $  96
// Virgin Galactic    39 One-way    $  37
// SpaceX             31 One-way    $  41
// Space Adventures   22 Round-trip $ 100
// Space Adventures   22 One-way    $  50
// Virgin Galactic    30 Round-trip $  84
// Virgin Galactic    24 Round-trip $  94
// Space Adventures   27 One-way    $  44
// Space Adventures   28 Round-trip $  86
// SpaceX             41 Round-trip $  72

const distanceToMars = 62100000

func genRandomSpeed() (randSpeed int) {
	minSpeed := 16
	maxSpeed := 30
	rand.Seed(time.Now().UTC().UnixNano())
	randSpeed = rand.Intn(maxSpeed-minSpeed) + minSpeed
	return
}

func roundTrip() (roundTrip bool) {
	return rand.Intn(2) == 1
}

func returnCarrier() (carrier string) {
	carriers := []string{
		"Virgin Galactic",
		"Space Adventures",
		"NASA",
		"Ian's Spacecapade",
	}
	index := rand.Intn(4)
	carrier = carriers[index]
	return carrier
}

func cost(days int) float64 {
	return math.Round(float64(days) * 0.8)
}

func main() {
	fmt.Println("Spaceline        Days Trip Type  Price")
	fmt.Println(strings.Repeat("=", 38))

	for count := 0; count < 10; count++ {
		speed := genRandomSpeed()
		days := distanceToMars / speed / 60 / 60 / 24
		carrier := returnCarrier()
		roundTrip := roundTrip()
		trip := "One-Way"

		if roundTrip {
			days = days * 2
			trip = "Round-Trip"
		}

		totalCost := cost(days)

		fmt.Printf("%-17s %3d %-10s $ %3g\n", carrier, days, trip, totalCost)
	}
}
