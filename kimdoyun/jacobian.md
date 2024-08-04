# jacobian 행렬

###### 김도윤 

### Go Code
```go
package main

import (
	"fmt"
	"math"
)

type MultivariableFunction interface {
	Evaluate(x, y float64) float64
}

type f1 struct{}

func (f f1) Evaluate(x, y float64) float64 {
	return x*x + y*y
}

type f2 struct{}

func (f f2) Evaluate(x, y float64) float64 {
	return x*y + 2*x - y
}

type f3 struct{}

func (f f3) Evaluate(x, y float64) float64 {
	return math.Pow(x, 3) - math.Pow(y, 3)
}

type Jacobian struct {
	functions []MultivariableFunction
}
func (j *Jacobian) Calculate(x, y float64) [][]float64 {
	delta := 1e-6
	rows := len(j.functions)
	jacob := make([][]float64, rows)
	for i, fn := range j.functions {
		jacob[i] = make([]float64, 2)
		jacob[i][0] = (fn.Evaluate(x+delta, y) - fn.Evaluate(x, y)) / delta
		jacob[i][1] = (fn.Evaluate(x, y+delta) - fn.Evaluate(x, y)) / delta
	}

	return jacob
}

func main() {
	x := 2.0
	y := 3.0
	functions := []MultivariableFunction{
		f1{},
		f2{},
		f3{},
	}
	jacobian := Jacobian{functions: functions}
	jacob := jacobian.Calculate(x, y)
	fmt.Printf("야코비안 행렬 (x = %.2f, y = %.2f):\n", x, y)
	for _, row := range jacob {
		for _, value := range row {
			fmt.Printf("%.4f ", value)
		}
		fmt.Println()
	}
}