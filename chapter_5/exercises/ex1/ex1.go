package main

import "fmt"

type person struct {
	first, last      string
	icecreamFlavours []string
}

func main() {
	p1 := person{
		first:            "james",
		last:             "bond",
		icecreamFlavours: []string{`Banana`, `Chocolat`},
	}
	p2 := person{
		first:            "miss",
		last:             "moneypenny",
		icecreamFlavours: []string{`Strawberry`, `Vanilla`},
	}
	for _, v := range p1.icecreamFlavours {
		fmt.Println(v)
	}
	for _, v := range p2.icecreamFlavours {
		fmt.Println(v)
	}
}
