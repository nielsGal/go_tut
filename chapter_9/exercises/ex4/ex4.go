package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mu sync.Mutex

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64

	const gs = 150
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			runtime.Gosched() //swap runtimes
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
