package main

import "fmt"

func main() {
	{
		x := "hello world"
		fmt.Println(x)
	}
	{
		x := "hello again"
		fmt.Println(x)
	}
	fmt.Println(incrementor()())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
