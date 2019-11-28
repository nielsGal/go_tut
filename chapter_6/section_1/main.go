package main

import "fmt"

func main() {
	foo()
	s := "james"
	bar(s)
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	s2, s3 := woowoo("james")
	fmt.Println(s2, s3)
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("hello", s)
}

func woo(s string) string {
	return fmt.Sprintf("hellow from woo %s \n", s)
}

func woowoo(s string) (string, string) {
	res1 := fmt.Sprintf("hello from woowoo %s ", s)
	res2 := fmt.Sprintf("hello again from woowoo %s", s)
	return res1, res2
}
