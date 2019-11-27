package main

import "fmt"

func main() {
	x := 42 //first time use walrus
	fmt.Println(x)
	x = 19 //second time use normal operator
	fmt.Println(x)
}
