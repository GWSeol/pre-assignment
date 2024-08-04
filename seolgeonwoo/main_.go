package main

import (
	"fmt"
)

func main() {
	f := FunctionImpl{
		domain: []int{2},
		funcValues: func(x []float64) []float64 {
			return []float64{
				x[0] * x[0], // 예시로 f(x) = x^2
			}
		},
	}

	point := Point{
		Coordinates: []float64{2.0},
	}

	diffMatrix := Differential(f, point)

	fmt.Println("Differential Matrix:")
	for _, row := range diffMatrix {
		fmt.Println(row)
	}
}
