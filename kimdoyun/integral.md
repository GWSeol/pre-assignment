## 2차원 함수의 적분 프로그램

김도운

```go
package main

import (
	"fmt"
)

// IntegralPossibility 인터페이스 정의
type IntegralPossibility interface {
	IntegralContinuity() bool // 적분 가능성 판별 함수: 연속인지 아닌지
}

// IntegralVisualizing 인터페이스 정의
type IntegralVisualizing interface {
	IntegralPossibility
	Integral(func(float64) float64) float64
}

// 예제 함수 정의
func function(x float64) float64 {
	return x * x // 예를 들어 함수 f(x) = x^2
}

// 적분 가능성 판별 함수
func IntegralContinuity() bool { // 적분 가능성 판별 함수: 연속인지 아닌지
	for x := -100.0; x <= 100.0; x++ { // [-100,100]에서 각 점에 해당하는 함수값을
		value := function(x)
		if value < -1000000000 || value > 1000000000 {
			return false
		}
	}
	return true
}

// 적분 계산 함수
func Integral(function func(float64) float64) float64 {
	alpha := 1.0 / 50000.0
	sum := 0.0
	for i := -100.0; i <= 100.0; i += alpha {
		sum += function(i) * alpha
	}
	return sum
}

// 디렉터 함수
func Director() string {
	if IntegralContinuity() {
		return fmt.Sprintf("Integral result: %.6f", Integral(function))
	} else {
		return "Integral not possible"
	}
}

func main() {
	// Director 함수를 호출하고 결과를 출력
	result := Director()
	fmt.Println(result)
}
-------------------------------------------------------
-출력결과
Integral result: 666666.666704 
-------------------------------------------------------
-실제결과
Intergral result: 666666.667(오차율: 약 0%)