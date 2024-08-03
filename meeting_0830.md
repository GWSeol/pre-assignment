# Meeting Notes for 8/3/2024

## Point

```go
type Point interface {
    Coordinate() []float64
    Dimension() int
}
```

```go
type Point []float64
```

* struct와 interface의 개념을 구분해야함: 언제 struct를 쓰고 언제 interface를 쓸 것인가
* struct로 구현해보자.

> ```go
> type Point []float64
> ```

* 아래와 같은 상황이 생기지 않도록 주의

```go
type SingleFunction struct {
    Interval [2]float64
    Eval func(float64) float64
}
type DoubleFunction struct {
    Rectangular [2][2]float64
    Eval func(float64, float64) float64
}
func SingleEvaluation(f SingleFunction, x Point) float64 {
    return f.Eval(x[0])
}
func DoubleEvaluation(f DoubleFunction, x Point) float64 {
    return f.Eval(x[0], x[1])
}
```

## Set

```go
type Set interface {
	containChecker(any) bool
}
```

```go
type Domain interface {
    Point
}
```

```go
type Set map[string]bool
```

* Set은 인터페이스로

> ```go
> type Set[T any] interface {
>     Check(T) bool
> }
> ```

* 아래는 참고용

```go
type Real float64

func (r Real) Check(x float64) bool {
    return true
}

var _ Set[float64] = Real(0)

type Interval interface {
    Set[Real]
    Lower() Real
    Upper() Real
}
```

## Function

```go

type Function struct {
	domain      Set
	codomain    Set
	maprelation func(any) any
}
```

```go
type MultiVarFunction struct {
    h float64 // 편미분 계산을 위한 작은 값
}
```

```go
type Function interface {
    Map(float64, float64) float64
    Domain
}
```

* Function은 struct

```go
type Function interface {
    // 내용은 알아서
}
```

* 아래는 참고용

```go
type Single struct {
    Interval // struct는 interface를 상속(?)할 수 있다.
    Eval func(float64) float64
}

type Itvl [2]float64

func (i Itvl) Lower() Real {
    return Real(i[0])
}
func (i Itvl) Upper() Real {
    return Real(i[1])
}

var s := Single{
    Interval: Itvl{0, 1},
    Eval: func(x float64) float64 {
        return x
    }
}
```

## Diff

```go
type PartialDifferential interface {
    Function1(x, y float64) float64
    Function2(x, y float64) float64
    Differential1(func(float64, float64) float64, string) float64
    Differential2(func(float64, float64) float64, string) float64
    Partial1() (float64, float64)
    Partial2() (float64, float64)
}
```

```go
type Diff(Function, float64, float64) float64
```

* Diff(미분계수)는 struct
* `Matrix`를 구현해야할 것