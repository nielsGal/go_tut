package main

import "fmt"

func main() {
	xi := []int{1, 23, 45, 56, 7, 8, 89, 9}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(x ...int) int {
	res := 0
	for _, v := range x {
		res += v
	}
	return res
}
func bar(x []int) int {
	res := 0
	for _, v := range x {
		res += v
	}
	return res
}
