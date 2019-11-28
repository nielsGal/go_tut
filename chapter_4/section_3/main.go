package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46}
	fmt.Println(x[1:3])              //take a slice of the slice using python like indexing
	for key, value := range x[1:3] { //the key is simply zero starting at the lower bound
		fmt.Println(key, value)
	}
	x = append(x, 6, 7, 8, 9) //this is a static function, not part of x
	for key := range x {
		fmt.Println(key)
	}
	y := []int{6, 7, 8}
	z := append(x, y...) //java script like operator
	// ...T variadic param  T... unpack
	fmt.Println(z)
}
