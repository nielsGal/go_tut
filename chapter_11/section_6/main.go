package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	_, err := squareRoot(-10)
	fmt.Printf("%T %v \n", err, err)
	if err != nil {
		log.Fatalln(err)
	}
}

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("square root of negative number")
	}
	return 42, nil
}
