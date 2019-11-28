package main

import "fmt"

func main() {
	x := 42
	y := 43
	g := x == y
	h := x >= y
	i := x <= y
	j := x != y
	k := x < y
	l := y < x
	fmt.Println(g, h, i, j, k, l)

}
