package main

import (
	"fmt"
)

type FunctionImpl struct {
	domain     []int
	funcValues func(x []float64) []float64
}

type Point struct {
	Coordinates []float64
}

func Differential(f FunctionImpl, p Point) [][]float64 {
	x := p.Coordinates
	funcValues := f.funcValues(x)

	// 더미 미분 행렬 생성 (예시)
	diffMatrix := make([][]float64, len(funcValues))
	for i := range funcValues {
		diffMatrix[i] = []float64{funcValues[i]}
	}

	return diffMatrix
}

func main() {
	// FunctionImpl 정의
	f := FunctionImpl{
		domain: []int{2},
		funcValues: func(x []float64) []float64 {
			return []float64{
				x[0] * x[0], // 예: f(x) = x^2
			}
		},
	}

	point := Point{
		Coordinates: []float64{5.0},
	}

	diffMatrix := Differential(f, point)

	fmt.Println("미분계수 행렬:")
	for _, row := range diffMatrix {
		fmt.Println(row)
	}
}
