package main

import "fmt"

//for vars the scope is package
var y = 42                                                 // some type is assigned even if not explicitly given
var z = ` a rar string gives " " explicitly "hello world"` // some type is assigned here string

func main() {
	fmt.Println(y)
	//y = z this is not possible
	fmt.Printf("%T\n", y) // we can print a type
	fmt.Println(y)
	fmt.Printf("%T\n", z)
	fmt.Println(z)
}
