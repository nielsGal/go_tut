package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 8}
	x := sum(xi...)
	fmt.Println(x)
	y := sum()
	fmt.Println(y)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
