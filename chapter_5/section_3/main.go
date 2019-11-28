package main

import "fmt"

func main() {
	x := struct {
		first, last string
		age         int
	}{
		first: "james",
		last:  "bond",
		age:   32,
	}
	fmt.Println(x)
}
