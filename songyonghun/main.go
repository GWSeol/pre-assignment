package main

import (
	"fmt"
)

// 다항식을 표현하기 위한 구조체
type Polynomial struct {
	Coefficients []float64
}

// 다항식의 미분을 계산하는 함수
func (p *Polynomial) Differentiate() *Polynomial {
	n := len(p.Coefficients)
	if n == 0 {
		return &Polynomial{Coefficients: []float64{}}
	}
	if n == 1 {
		return &Polynomial{Coefficients: []float64{0}}
	}

	// 미분한 다항식의 계수를 저장할 슬라이스
	derivedCoefficients := make([]float64, n-1)

	// 각 항을 미분
	for i := 1; i < n; i++ {
		derivedCoefficients[i-1] = float64(i) * p.Coefficients[i]
	}

	return &Polynomial{Coefficients: derivedCoefficients}
}

// 다항식을 문자열로 표현하는 함수
func (p *Polynomial) String() string {
	str := ""
	for i := len(p.Coefficients) - 1; i >= 0; i-- {
		coeff := p.Coefficients[i]
		if coeff == 0 {
			continue
		}
		if coeff > 0 && i != len(p.Coefficients)-1 {
			str += "+"
		}
		str += fmt.Sprintf("%.2fx^%d", coeff, i)
	}
	return str
}

func main() {
	// 예시 다항식: 3x^3 + 2x^2 + 5x + 7
	p := &Polynomial{Coefficients: []float64{7, 5, 2, 3}}
	fmt.Println("원래 다항식:", p)

	// 다항식 미분
	derivedP := p.Differentiate()
	fmt.Println("미분한 다항식:", derivedP)
}
