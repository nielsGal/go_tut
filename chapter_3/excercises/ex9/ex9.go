package main

import "fmt"

func main() {
	favSport := "snowboarding"
	switch favSport {
	case "true": // you cannot use a non string in this match
		fmt.Println("true matches")
	case "snowboarding":
		fmt.Println("just like scala")
	}
}
