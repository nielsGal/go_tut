package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "chocolate", "martini"}
	mp := []string{"Miss", "MoneyPenny", "Strawberry", "Hazelnut"}
	fmt.Println(jb)
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
	m := map[string]int{
		"james":      42,
		"moneypenny": 27,
	}
	fmt.Println(m["james"])  //value in map
	fmt.Println(m["james2"]) //value not in map return the 0
	if v, ok := m["james2"]; ok {
		println(v)
	}
	if v, ok := m["james"]; ok {
		println(v)
	}
	m["james2"] = 34
	if v, ok := m["james2"]; ok { //now we do have the value
		println(v)
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
	delete(m, "james")
	if v, ok := m["james"]; ok { // now james is gone
		println(v)
	} else {
		println(ok)
	}
}
