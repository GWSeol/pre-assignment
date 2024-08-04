package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y, Z float64
}

func Distance(p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	dz := p2.Z - p1.Z

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func main() {
	point1 := Point{X: 1.0, Y: 2.0, Z: 2.0}
	point2 := Point{X: 4.0, Y: 6.0, Z: 3.0}

	distance := Distance(point1, point2)

	fmt.Printf("두 점사이의 거리: %.2f\n", distance)
}
