package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "hello world\n")
	io.WriteString(os.Stdout, "hello world\n")
}
