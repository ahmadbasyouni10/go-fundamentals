package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	var float32Slice = []float32{1.1, 2.2, 3.3}
	var float64Slice = []float64{1.1, 2.2, 3.3}
	fmt.Println(sumSlice[int](intSlice))
	fmt.Println(sumSlice(float32Slice))
	fmt.Println(sumSlice[float64](float64Slice))
	var emptySlice []int

	fmt.Println(isEmpty(emptySlice))
	fmt.Println(isEmpty[float32](float32Slice))
	fmt.Println(isEmpty[float64](float64Slice))
	fmt.Println(isEmpty[int](emptySlice))
}

// using T any wouldnt work since we are using + operator which is not defined for all types
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// to not repeat yourself and make more functions with same signature we can use generics
