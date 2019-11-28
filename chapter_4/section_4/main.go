package main

import "fmt"

func main() {
	x := []int{1, 23, 45, 7, 8, 9, 10}
	x = append(x[:2], x[4:]...) //delete numbers x[3] x[4]
	fmt.Println(x)
	y := make([]int, 10, 12)
	fmt.Printf("%T %v\n", y, y)
	//y[11] = 40   this will error
	y = append(y, 12)
	fmt.Println(cap(y))
	y = append(y, 12)
	fmt.Println(cap(y))
	y = append(y, 12) //this one will double the cap
	fmt.Println(cap(y))

	fmt.Printf("%T %v\n", y, y)
}
