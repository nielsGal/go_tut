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
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.icecreamFlavours {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
