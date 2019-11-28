package main

import "fmt"

//do not use real arrays use slices instead!

func main() {
	var x [5]int   //an array of size 5
	var y [6]int   // another arrat of size 6
	fmt.Println(x) //we can just print these unlike java
	x[3] = 42      //0 based indexing like normal
	fmt.Println(x)
	fmt.Printf("%T %T \n", x, y) //the types are disticnt [5]int vs [6]int
}
