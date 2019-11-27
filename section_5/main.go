package main

import "fmt"

var a int = 25 //a global cannot use the walrus (:=)

func main() {
	var x int = 42 //var identifier type,
	fmt.Println(x)
	var y, z int = 22, 23 //mutli assignment
	fmt.Println(y, " ", z)
	fmt.Println(a) // this one is global, it is not checked for use.
}
