package main

import (
    "fmt"
)

type Set map[string]bool

func NewSet() Set {
    return make(Set)
}

func (s Set) Add(element string) {
    s[element] = true
}

func (s Set) Remove(element string) {
    delete(s, element)
}

func (s Set) Contains(element string) bool {
    _, exists := s[element]
    return exists
}

func (s Set) Size() int {
    return len(s)
}

func (s Set) Union(other Set) Set {
    result := NewSet()
    for element := range s {
        result.Add(element)
    }
    for element := range other {
        result.Add(element)
    }
    return result
}

func (s Set) Intersection(other Set) Set {
    result := NewSet()
    for element := range s {
        if other.Contains(element) {
            result.Add(element)
        }
    }
    return result
}

func (s Set) Difference(other Set) Set {
    result := NewSet()
    for element := range s {
        if !other.Contains(element) {
            result.Add(element)
        }
    }
    return result
}

func main() {
    
}