package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4, 5, 6, 8)
	fmt.Println(x)
}

func sum(x ...int) int {
	fmt.Printf("%T , %v \n", x, x)

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
