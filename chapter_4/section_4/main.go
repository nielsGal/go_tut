package main

import "fmt"

func main() {
	x := []int{1, 23, 45, 7, 8, 9, 10}
	x = append(x[:2], x[4:]...) //delete numbers x[3] x[4]
	fmt.Println(x)
}
