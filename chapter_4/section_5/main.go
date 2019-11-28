package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "chocolate", "martini"}
	mp := []string{"Miss", "MoneyPenny", "Strawberry", "Hazelnut"}
	fmt.Println(jb)
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
