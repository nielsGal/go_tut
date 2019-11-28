package main

import "fmt"

func main() {
	//x := type{values}
	x := []int{12, 23, 34, 45, 56}
	fmt.Printf("%T %v\n", x, x) //this does not have type [5]int but type []int
	for _, value := range x {   //throw away key pring value
		fmt.Println(value)
	}
	for key, _ := range x {
		fmt.Println(key)
	}
	for key := range x { //this is preferable
		fmt.Println(key)
	}
}
