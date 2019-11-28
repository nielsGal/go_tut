package main

import "fmt"

//do not use real arrays use slices instead!

func main() {
	var x [5]int   //an array of size 5
	fmt.Println(x) //we can just print these unlike java
	x[3] = 42      //0 based indexing like normal
	fmt.Println(x)
}
