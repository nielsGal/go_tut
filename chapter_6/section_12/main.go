package main

import "fmt"

func main() {
	fmt.Println(fib(15))
}

func fib(x int) int {
	if x <= 1 {
		return 1
	} else {
		return fib(x-1) + fib(x-2)
	}
}
