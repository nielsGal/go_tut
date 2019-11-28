package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	z := [][]string{x, y}
	for k, v := range z {
		fmt.Println(k, v)
		for key, value := range v {
			fmt.Println(key, value)
		}
	}

}
