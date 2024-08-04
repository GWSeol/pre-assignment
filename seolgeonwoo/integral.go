//적분 구간 1~0까지의 x에대한 일변수함수 정적분

package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func parseFunction(exprStr string) (func(float64) float64, error) {
	program, err := expr.Compile(exprStr)
	if err != nil {
		return nil, err
	}
	return func(x float64) float64 {
		result, _ := expr.Run(program, map[string]interface{}{"x": x})
		return result.(float64)
	}, nil
}

func rect(a, b float64, n int, f func(float64) float64) float64 {
	h := (b - a) / float64(n)
	sum := (f(a) + f(b)) / 2.0
	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		sum += f(x)
	}
	return sum * h
}

func main() {
	var input string
	fmt.Println("x에 대한 함수를 입력하세요(ex_ x^2+3*x+9):")
	fmt.Scanln(&input)

	// 문자열에서 함수로 변환
	f, err := parseFunction(input)
	if err != nil {
		fmt.Println("다시 입력해주세요:", err)
		return
	}

	a := 0.0  // 적분 구간의 시작
	b := 1.0  // 적분 구간의 끝
	n := 1000 // 

	integral := rect(a, b, n, f)
	fmt.Printf("적분 값은 %.2f to %.2f 약 %.5f\n입니다.", a, b, integral)
}
