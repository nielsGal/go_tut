package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(f)
	_, err = os.Open("myfile.txt")
	if err != nil {
		log.Println(err)
	}
	defer foo()
	_, err = os.Open("myfile.txt")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("succes")
}

func foo() {
	fmt.Println("defferred functions don't run on os.exit")
}
