package main

import (
	"fmt"
	"sync"
)

var raceInt int
var wg sync.WaitGroup

func main() {
	wg.Add(5)
	go race()
	go race()
	go race()
	go race()
	go race()
	wg.Wait()
	fmt.Println(raceInt)
}

func race() {
	for i := 0; i < 10000; i++ {
		raceInt++
	}
	wg.Done()
}
