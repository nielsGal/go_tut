package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	//%T type, %d base 10 %b binary %o octo, %x hex, %#x hex with 0x
	fmt.Printf("%T %d %b %o %x %#x\n", y, y, y, y, y, y)
}
