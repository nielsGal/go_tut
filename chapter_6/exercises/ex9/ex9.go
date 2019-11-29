package main

func main() {
	bar(foo, "bar")
}
func foo(s string) {
	println("foo", s)
}
func bar(f func(s string), s string) {
	f(s)
}
