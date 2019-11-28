package main

import "fmt"

func main() {
	x := 0
	y := 1994
	for {
		if x > 25 {
			break
		}
		fmt.Println(y + x)
		x++
	}
}
