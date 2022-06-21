package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{
		"first":  10,
		"second": 20,
	}

	floats := map[string]float64{
		"first":  35.01,
		"second": 20.05,
	}

	fmt.Printf("Non-generics Sum: %v and %v\n", SumInts(ints), SumFloats(floats))
}
