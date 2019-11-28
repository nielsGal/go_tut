package main

import "fmt"

func main() {
	if z := false; z {

	} else if z {
		fmt.Println(z) //z is still accesible here
	} else {
		fmt.Println(z) //z is still accesible here
	}
}
