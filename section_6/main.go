package main

import "fmt"

var y = 42            // some type is assigned even if not explicitly given
var z = "hello world" // some type is assigned here string

func main() {
	fmt.Println(y)
	//y = z this is not possible
	fmt.Printf("%T\n", y) // we can print a type
	fmt.Printf("%T\n", z)
}
