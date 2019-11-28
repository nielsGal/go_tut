package main

import "fmt"

var x = 42
var y = "james bond"
var z = true

func main() {
	s := fmt.Sprintf("%d %s %t", x, y, z)
	fmt.Println(s)
}
