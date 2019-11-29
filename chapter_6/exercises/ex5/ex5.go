package main

const (
	pi = 3.14
)

type circle struct {
	radius float64
}
type square struct {
	side float64
}
type shape interface {
	area() float64
}

func main() {
	c := circle{
		radius: 1.5,
	}
	s := square{
		side: 2.5,
	}
	info(c)
	info(s)
}

func (c circle) area() float64 {
	return pi * c.radius * c.radius
}
func (s square) area() float64 {
	return s.side * s.side
}
func info(s shape) {
	switch s.(type) {
	case square:
		println("this is a square with area", s.area())
	case circle:
		println("this is a circel with area", s.area())
	}
}
