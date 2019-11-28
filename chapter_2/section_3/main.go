package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 42
	y := 42.333
	fmt.Printf("%T %T %v %v\n", x, y, x, y)
	fmt.Println(float64(x) == y)
	var z int8 = -128 // -129 errs
	fmt.Println(z)
	var a int
	var b uint
	var c uintptr
	fmt.Printf("%v %v %v\n", a, b, c) // these change 32 vs 64 bit depeding on the machine
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}
