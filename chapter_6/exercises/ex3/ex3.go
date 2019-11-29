package main

func main() {
	defer foo()
	bar()
}
func foo() {
	println("hello from foo ")
}
func bar() {
	println("hello from bar")
}
