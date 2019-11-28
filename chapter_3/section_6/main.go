package main

import "fmt"

func main() {
	x := true
	y := false
	if x {
		fmt.Println(x)
	}
	if !y {
		fmt.Println(y)
	}
	if 2 == 2 { // the go fm changes (2==2) to 2 == 2
		fmt.Println(x)
	}
	if z := true; z { // z is now scoped to this if
		fmt.Println(z)
	}
	//fmt.Println(z) you cannot use z here.
}
