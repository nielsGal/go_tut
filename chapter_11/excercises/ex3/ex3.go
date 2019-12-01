package main

import "fmt"

type customErr struct {
}

func main() {
	errmsg := customErr{}
	foo(errmsg)
}
func foo(e error) {
	fmt.Println(e)
}

func (e customErr) Error() string {
	return "an error"
}
