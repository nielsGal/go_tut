package main

import "fmt"

func main() {
	x := "a string"
	y := `"a raw string"`
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	bs := []byte(x)
	fmt.Printf("%v %s %T\n", bs, bs, bs)

	for i := 0; i < len(x); i++ {
		fmt.Printf("%#U", x[i])
	}
	for i, v := range x {
		println(i, v)
	}
}
