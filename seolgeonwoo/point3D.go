package main

import (
	"fmt"
	"math"
)

// Point3D 구조체는 3차원 좌표를 나타냅니다.
type Point3D struct {
	X, Y, Z float64
}

// Distance 함수는 두 점 사이의 3차원 거리를 계산합니다.
func Distance(p1, p2 Point3D) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	dz := p2.Z - p1.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func main() {
	var x1, y1, z1, x2, y2, z2 float64

	// 사용자로부터 첫 번째 점의 좌표를 입력받습니다.
	fmt.Print("좌표입력 x1 y1 z1: ")
	_, err := fmt.Scanf("%f %f %f", &x1, &y1, &z1)
	if err != nil {
		fmt.Println("에러:", err)
		return
	}

	// 사용자로부터 두 번째 점의 좌표를 입력받습니다.
	fmt.Print("좌표입력 x2 y2 z2 : ")
	_, err = fmt.Scanf("%f %f %f", &x2, &y2, &z2)
	if err != nil {
		fmt.Println("에러:", err)
		return
	}

	// 두 점을 정의합니다.
	point1 := Point3D{X: x1, Y: y1, Z: z1}
	point2 := Point3D{X: x2, Y: y2, Z: z2}

	// 두 점 사이의 거리를 계산합니다.
	distance := Distance(point1, point2)

	// 결과를 출력합니다.
	fmt.Printf("두 점 사이의 거리: %.2f\n", distance)
}
