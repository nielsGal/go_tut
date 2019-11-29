package main

import "fmt"

func main() {

	func() {
		fmt.Println("hello from anon")
	}()
	func(x int) {
		fmt.Println("hello from anon", x)
	}(42)
}
