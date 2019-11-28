package main

import "fmt"

func main() {
	const (
		x     = 50
		y int = 60
	)
	fmt.Printf("%T %T \n %v %v", x, y, x, y)
}
