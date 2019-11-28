package main

import (
	"fmt" //multiple possible imports, single can be inline
)

var a int = 25 //a global cannot use the walrus (:=)
var b int      // some zero value exists for all types

func main() {
	var x int = 42 //var identifier type,
	fmt.Println(x)
	var y, z int = 22, 23 //mutli assignment
	fmt.Println(y, " ", z)
	fmt.Println(a) // this one is global, it is not checked for use.
	fmt.Println(b)
}
