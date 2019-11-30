package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("architecture", runtime.GOARCH)
	fmt.Println("num CPU", runtime.NumCPU())
	fmt.Println("num go routines", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("num go routines", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("bar", i)
	}

	wg.Done()
}
