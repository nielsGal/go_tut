package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First, Last string
	Age         int
}

func main() {
	p1 := person{
		First: "james",
		Last:  "bond",
		Age:   32,
	}
	p2 := person{
		First: "miss",
		Last:  "moneypenny",
		Age:   27,
	}
	people := []person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
	fmt.Println(string(bs))
}
