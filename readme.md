### 2024 PreURP 

### Go 언어를 이용한 미분기하학적 개념의 구조적 접근
#### Systematic approach to differential geometric concepts using Go programming language

### 사전 과제

> "It is not who you are underneath, it is what you do that defines you."
> *Batman Begins (2005)*

***기하학의 추상적 개념을 고 언어의 인터페이스 타입으로 정의하자***

#### 추상적 개념의 예시
1. **집합**은 원소의 포함여부와 원소의 추출과 삽입이 가능한 대상이다. **유한집합**은 원소의 **조합**을 만들 수 있다.
2. **점**은 공간에서 위치를 나타내는 대상이다. **거리**는 두 점 사이의 관계를 나타내는 대상이다.
3. **함수**는 입력과 출력 사이의 관계를 나타내는 대상이다. 몇몇 함수는 **그래프**를 그릴 수 있다. 
4. **미분**은 입력값의 변화에 따른 출력값의 변화를 **선형근사**한다. 

### 사전 과제 목표

#### 1. Go 언어 설치
* Go 다운로드 및 설치
  * https://golang.org/dl/
* Visual Studio 다운로드 및 설치
  * https://code.visualstudio.com
  * Go 확장 프로그램 설치
    ![](ext_go.png)

#### 2. Go 언어 기본 문법 익히기
* Go 언어 기본 문법 학습
  * https://tour.golang.org/welcome/1
* Go 언어 특징 학습
  * https://gobyexample.com/multiple-return-values
  * https://gobyexample.com/variadic-functions
  * https://gobyexample.com/closures
  * https://gobyexample.com/recursion
  * https://gobyexample.com/pointers
  * https://gobyexample.com/structs
  * https://gobyexample.com/methods
  * https://gobyexample.com/interfaces
  * https://gobyexample.com/struct-embedding
  * https://gobyexample.com/generics
  * https://gobyexample.com/goroutines
  * https://gobyexample.com/channels
  * https://gobyexample.com/channel-directions
  * https://gobyexample.com/range-over-channels
  * https://gobyexample.com/panic
  * https://gobyexample.com/defer

#### 3. 집합(set) 구현
> 유한집합(finite set)은 그 원소 중 $r$개를 (중복없이) 선택하여 조합(combination)을 만들 수 있다. 모든 조합을 생성하기 위해 `Set` 타입이 가져야 할 기능(method)을 정의하라.
* `Set` 타입 정의
  * 정수의 집합 예: `[]int`, `map[int]struct{}`, etc.
  * 실수의 집합 예: `[]float64`, `map[float64]struct{}`, etc.
* `CombinationSet` 메소드 정의
  > `Set` 타입의 원소의 조합을 반환하는 함수.
  ```go
  func CombinationSet(Set, int) []Set
  ```
* `S` 인터페이스 타입 정의
  > `Set` 타입이 가져야 할 메소드를 정의한 인터페이스.
  ```go
  type S interface {
    // write your methods 
  }
  ```
* `Combination` 함수 구현
  > `S` 타입의 원소의 조합을 반환하는 함수
  ```go
  func Combination(S, int) <-chan S 
  ```

#### 4. 점(point) 구현
> 점(point)은 $n$차원 공간에서 위치를 나타내는 개념이다. 점의 위치를 나타내기 위해 `Point` 타입이 가져야 할 기능(method)을 정의하라.
* `Point` 타입 정의
    * 예: `[]float64`, `[2]float64`, `[3]float64`, etc.
* `P` 인터페이스 타입 정의
    > `Point` 타입이 가져야 할 메소드를 정의한 인터페이스.
    ```go
    type P interface {
        // write your methods  
    }
    ```
* `Distance` 함수 구현
    > 두 점 사이의 거리를 반환하는 함수
    ```go
    func Distance(P, P) float64
    ```

#### 5. 함수(function) 구현
> 함수(function)은 입력과 출력 사이의 관계를 나타내는 개념이다. 함수의 관계를 나타내기 위해 `Function` 타입이 가져야 할 기능(method)을 정의하라.
* `Function` 타입 정의
    * 예: `func(float64) float64`, `func(float64, float64) float64`, `func(P) P`, etc.
* `F` 인터페이스 타입 정의
    > 함수가 가져야 할 메소드를 정의한 인터페이스.
    ```go
    type F interface {
        // write your methods  
    }
    ```
* `Graph` 함수 구현
    > 함수의 그래프를 그리는 함수
    ```go
    func Graph(F) 
    ```
