package main

import "fmt"

func main() {
	fmt.Println("bar")
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}
func foo() {
	n, err := fmt.Println("foo")
	fmt.Println(n, err)
}
