package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	c1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "dark cyan",
		},
		fourWheel: true,
	}
	c2 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "bright red",
		},
		luxury: true,
	}
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c1.doors)
	fmt.Println(c2.doors)
}
