package main

import "fmt"

func main() {
	x := 2
	y := 5
	for i := 0; i < 100; i++ {
		switch {
		case i%x == 0:
			fmt.Print("fizz")
			fallthrough
		case i%y == 0:
			fmt.Print("buzz")
		default:
			fmt.Print(" ")
		}
	}
	s := "bond"
	switch s {
	case "bond", "james":
		fmt.Println("james bond")
	}
	switch "james" {
	case "bond", "james":
		fmt.Println("james bond")
	}
}
