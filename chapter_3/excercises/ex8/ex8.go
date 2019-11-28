package main

import "fmt"

func main() {
	s := "a string to match"
	switch {
	case s == "a string to match":
		fmt.Println("we matched the string")
	}
}
