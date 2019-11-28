package main

import "fmt"

func main() {
	a := 2
	b := 1500
	for a < b { //this is basically a while
		a *= 2
	}
	fmt.Println(a, b)
	c := 1
	d := 5
	for ; c < d; c++ { //the init i optional the ; is not
		b *= 2
	}
	fmt.Println(b, c, d)
	e := 0
	for { //forever use for webservers
		e++
		if e > 10 {
			break
		}
	}
	fmt.Println(e)
}
