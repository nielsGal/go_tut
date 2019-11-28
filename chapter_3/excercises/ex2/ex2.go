package main

import "fmt"

func main() {
	y := 91
	for x := 65; x < y; x++ {
		fmt.Printf("%d \n\t %U %c\n\t %U %c\n\t %U %c\n", x, x, x, x, x, x, x)
	}
}
