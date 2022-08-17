package main

import (
	"fmt"
)

func MeterToCentimeter(distanceInMeter int) int {

	distanceInCm := distanceInMeter * 100
	return distanceInCm
}

func main() {

	var distanceInMeter int
	fmt.Print("Type the distance in meters: ")
	fmt.Scanf("%d\n", &distanceInMeter)

	distanceInCm := MeterToCentimeter(distanceInMeter)
	fmt.Println(fmt.Sprintf("%vm is equal to %vcm", distanceInMeter, distanceInCm))
}
