package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	minf := min[float64]
	fmt.Println(minf(2.2, 3.3))
	fooDemo()
}

type Value interface {
	int | float64
}

func min[T Value](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func foo[S ~[]E, E any](s S) {
	fmt.Println(s)
}

type myStr string

func fooDemo() {
	str := myStr("hello")
	strs := []myStr{str, "world"}
	foo(strs)
}

// Scale 返回切片中每个元素都乘c的副本切片
func Scale[E constraints.Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}
