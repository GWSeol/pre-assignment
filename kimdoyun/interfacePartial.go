package main

import (
	"fmt"
	"math"
)

// 다변수 함수 인터페이스
type MultivariableFunction interface {
	Evaluate(x, y float64) float64
}

// f1 구조체 정의
type f1 struct{}

func (f f1) Evaluate(x, y float64) float64 {
	return x*x + y*y
}

// f2 구조체 정의
type f2 struct{}

func (f f2) Evaluate(x, y float64) float64 {
	return x*y + 2*x - y
}

// f3 구조체 정의
type f3 struct{}

func (f f3) Evaluate(x, y float64) float64 {
	return math.Pow(x, 3) - math.Pow(y, 3)
}

// Jacobian 구조체 정의
type Jacobian struct {
	functions []MultivariableFunction
}

// 야코비안 행렬 계산 메서드
func (j *Jacobian) Calculate(x, y float64) [][]float64 {
	delta := 1e-6 // 작은 값
	rows := len(j.functions)
	jacob := make([][]float64, rows)

	for i, fn := range j.functions {
		jacob[i] = make([]float64, 2) // 두 개의 편미분을 저장
		// ∂f/∂x
		jacob[i][0] = (fn.Evaluate(x+delta, y) - fn.Evaluate(x, y)) / delta
		// ∂f/∂y
		jacob[i][1] = (fn.Evaluate(x, y+delta) - fn.Evaluate(x, y)) / delta
	}

	return jacob
}

func main() {
	x := 2.0 // 미분할 점 x
	y := 3.0 // 미분할 점 y

	functions := []MultivariableFunction{
		f1{},
		f2{},
		f3{},
	}

	jacobian := Jacobian{functions: functions}
	jacob := jacobian.Calculate(x, y)

	// 야코비안 행렬 출력
	fmt.Printf("야코비안 행렬 (x = %.2f, y = %.2f):\n", x, y)
	for _, row := range jacob {
		for _, value := range row {
			fmt.Printf("%.4f ", value)
		}
		fmt.Println()
	}
}
