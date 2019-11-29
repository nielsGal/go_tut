package main

import "fmt"

func main() {
	a := 1
	println(a)
	println(&a)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", &a)

	//var b int = &a this gives type error
	var b *int = &a
	fmt.Printf("%T %v\n", b, *b)
	*b++
	fmt.Println(a)
	fmt.Println(*&b)
}
