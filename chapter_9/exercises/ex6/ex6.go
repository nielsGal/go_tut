package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOOS", runtime.GOOS)
	fmt.Println("GOOARCH", runtime.GOARCH)

}
