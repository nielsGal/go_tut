package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("architecture", runtime.GOARCH)
	fmt.Println("num CPU", runtime.NumCPU())
	fmt.Println("num go routines", runtime.NumGoroutine())
	go foo()
	go bar()
	fmt.Println("num go routines", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}
