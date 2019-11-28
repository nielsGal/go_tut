package main

import "fmt"

func main() {
	if s := "some special syntax"; s == "some special" {

	} else if s == "some special syntax" {
		fmt.Println(`go can use some special syntax`)
	}
}
