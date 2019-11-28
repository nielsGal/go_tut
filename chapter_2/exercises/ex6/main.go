package main

import "fmt"

func main() {
	const (
		y1 = 2019 + iota
		y2 = 2019 + iota
		y3 = 2019 + iota
		y4 = 2019 + iota
	)
	fmt.Println(y1, y2, y3, y4)
}
