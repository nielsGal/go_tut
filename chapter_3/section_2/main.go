package main

import "fmt"

func main() {
	for i := 3; i < 5; i++ {
		for j := 5; j < 7; j++ {
			fmt.Println(i*j, i, j)
		}
	}
}
