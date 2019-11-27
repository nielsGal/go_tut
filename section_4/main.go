package main

import "fmt"

func main() {
	x := 42 //first time use walrus
	fmt.Println(x)
	x = 18 //second time use normal operator
	fmt.Println(x)
	y := x + 24 // walrus can contain vars
	fmt.Println(y)
}
