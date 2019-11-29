package main

import "fmt"

func main() {
	ii := []int{5, 6, 7, 8, 9, 10, 12, 14}
	fmt.Println(even(sum, ii...))
	fmt.Println(odd(sum, ii...))
	fmt.Println(sum(ii...))
}

func sum(x ...int) int {
	res := 0
	for _, v := range x {
		res += v
	}
	return res
}
func even(f func(xi ...int) int, x ...int) int {
	var yi []int
	for _, v := range x {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return sum(yi...)
}

func odd(f func(xi ...int) int, x ...int) int {
	var yi []int
	for _, v := range x {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return sum(yi...)
}
