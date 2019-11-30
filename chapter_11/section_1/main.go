package main

import "fmt"

func main() {
	n, err := fmt.Println("hello")
	if err != nil {
		fmt.Println("the error", err)
	}
	fmt.Println(n)
}
