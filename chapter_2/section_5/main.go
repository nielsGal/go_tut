package main

import "fmt"

func main() {
	const (
		a = 42
		b = 42.78
		c = "james bond"
	) //leavgin out type gives compiler flexibility to assign type
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	const (
		d = iota
		e = iota
		f = iota
	)
	fmt.Printf("%T %T %T\n", d, e, f)
	fmt.Println(d, e, f)
}
