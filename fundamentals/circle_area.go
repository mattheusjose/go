package main

import (
	"fmt"
)

func CircleArea(radius int) float32 {
	const PI = 3.1415
	var area float32 = PI * float32(radius*radius)

	return area
}

func main() {

	var radius int
	fmt.Print("Circle radius: ")
	fmt.Scanf("%d\n", &radius)

	circleArea := CircleArea(radius)

	fmt.Println(fmt.Sprintf("The Circle with %v radius has %v of area", radius, circleArea))

}
