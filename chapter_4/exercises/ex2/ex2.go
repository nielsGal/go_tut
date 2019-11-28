package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 7, 8, 9, 0}
	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Printf("%T \n", x)
}
