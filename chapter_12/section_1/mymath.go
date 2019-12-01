//Package mymath my math package
package mymath

import "fmt"

//func test tests sum
func test() {
	fmt.Println(6 == sum([]int{1, 2, 3}...))
}

//sum adds ints
func sum(x ...int) int {
	res := 0
	for _, v := range x {
		res += v
	}
	return res
}
