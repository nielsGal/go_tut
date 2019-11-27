package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	//%T type, %d base 10 %b binary %o octal, %x hex, %#x hex with 0x
	x := fmt.Sprintf("%T %d %b %o %x %#x\n", y, y, y, y, y, y)
	fmt.Print(x)
	fmt.Printf("%v \n", y) //the default value always
}
