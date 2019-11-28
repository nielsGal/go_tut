package main

import "fmt"

func main() {
	x1 := struct {
		hostname string
		ipaddr   []uint8
		protno   int
	}{
		hostname: "localhost",
		ipaddr:   []uint8{127, 0, 0, 1},
		protno:   80,
	}
	fmt.Println(x1)
}
