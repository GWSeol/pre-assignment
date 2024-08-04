package main

func Differential(f Function, point Point) [][]float64 {
	funcValueAtPoint := f.Apply(point.Coordinates)

	epsilon := 1e-6 //(10^(-6))

	numVars := len(point.Coordinates)
	numFuncs := len(funcValueAtPoint)

	diffMatrix := make([][]float64, numFuncs)
	for i := range diffMatrix {
		diffMatrix[i] = make([]float64, numVars)
	}

	// 각 변수에 대해 편미분을 계산
	for i := 0; i < numVars; i++ {
		perturbedPoint := make([]float64, numVars)
		copy(perturbedPoint, point.Coordinates)

		perturbedPoint[i] += epsilon
		perturbedFuncValue := f.Apply(perturbedPoint)

		for j := 0; j < numFuncs; j++ {
			diffMatrix[j][i] = (perturbedFuncValue[j] - funcValueAtPoint[j]) / epsilon
		}
	}

	return diffMatrix
}
