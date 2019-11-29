package main

func main() {
	f := foo()
	println(f())
}

func foo() func() int {
	return func() int {
		return 5
	}
}
