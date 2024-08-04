package main

type Domain interface {
	GetDomain() []int
}

type Map interface {
	Apply(x []float64) []float64
}

type Function interface {
	Domain
	Map
}

type FunctionImpl struct {
	domain     []int
	funcValues func([]float64) []float64
}

func (f FunctionImpl) GetDomain() []int {
	return f.domain
}

func (f FunctionImpl) Apply(x []float64) []float64 {
	return f.funcValues(x)
}
