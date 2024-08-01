# 중간 점검

## 1. 연구 진행 방법

* 개개인이 서로 다른 module 및 package를 구현할 것인지, 아니면
* 함께 하나의 module 과 package를 생성할 것인지 결정

## 2. 필수 type 및 func 구현

### 2-1. type `Point`

```go
type Point interface {
    // 메소드 정의
}
```

### 2-1. type `Function`

```go
type Function interface {
    Domain
    Map
}

type Domain interface {
    // 메소드 정의
}

type Map interface {
    // 메소드 정의
}
```

### 2-2. func `Differential`

> [미분계수] 점 $\mathbf a \in \mathbf R^m$에서의 함수 $\mathbf f:\mathbf R^m \to \mathbf R^n$의 미분계수(differential)는 다음과 같은 $n\times m$ 행렬(matrix)이다.
> $$ D\mathbf f(\mathbf a) = \begin{bmatrix} \frac{\partial f_1}{\partial x_1}(\mathbf a) & \cdots & \frac{\partial f_1}{\partial x_m}(\mathbf a) \\ \vdots & \ddots & \vdots \\ \frac{\partial f_n}{\partial x_1}(\mathbf a) & \cdots & \frac{\partial f_n}{\partial x_m}(\mathbf a)\end{bmatrix}$$

행렬은 선형변환(linear transformation)으로서 함수의 일종이다. 따라서 미분계수는 함수(Function) $\mathbf f$와 점(Point) $\mathbf a$로 (선형)함수를 반환하는 함수로 구현할 수 있다.

```go
func Differential(Function, Point) Function {
    // 내용
}
```

단, 함수의 미분이 항상 존재하는 것은 아니다. 만약 편미분이 존재하지 않거나, 미분가능성의 조건을 만족하지 않으면 위 함수는 `nil`을 반환해야 한다.

### 2-3. func `Integral`

> [함수의 적분] 함수의 적분은 정의역 $D$의 각 분할된 영역(region)의 크기 $\Delta_{\mathbf x}$에 함숫값 $f(\mathbf x)$을 곱한 것을 더한 것으로 근사된다. 여기서 $\mathbf x$는 영역을 분할하는 격자(lattice) $\Lambda(D)$의 점이다. 분할의 갯수 $N$이 무한히 커질때, 그 합이 수렴하면 이를 적분(integral)이라 한다.
> $$\int_D f = \lim_{N\to\infty} \sum_{\mathbf x\in \Lambda(D)} f(\mathbf x) \Delta_{\mathbf x}$$

함수의 적분은 정의역의 차원(dimension)에 관계없이 같은 원리로 계산할 수 있다. 심지어 적분은 함수가 연속이 아니더라도 계산할 수 있다. 그러나 모든 함수가 적분가능한 것은 아니다. 아래의 `Integral` 함수는 적분가능한 경우에만 적분값을 반환해야 한다.

```go
func Integral(Function) float64 {
    // 내용
}
```

## 3. 앞으로의 계획 세우기

### 3-1. 구현할 타입과 함수 예시

* 벡터장 `VectorField`
* 곡선 `Curve`과 곡면 `Surface`
* 곡선 또는 곡면 위에서의 벡터장의 적분
    * 이를 위해 향 `Orientation`의 개념이 필요함
* 미분형식 `Form`
    * 미분형식의 적분
* 다양체 `Manifold`
    * 다양체의 정의 읽어보기
    * 다양체의 부피(volume)를 어떻게 정의하는가?

### 3-2. 결과물의 형태

* 구현한 타입을 해당 수학 개념과 함께 설명하는 메뉴얼 형식의 문서 
* 각자 책을 써 보는 것도 좋은 방법 - 정확한 수학적 개념을 어떻게 실용적으로 전달할 것인가?
* 각자 생성한 코드에 대한 비판적 분석을 글로 정리하는 것도 좋음
