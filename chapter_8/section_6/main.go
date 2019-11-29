package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "correcthorsebatterystaple"
	fmt.Println(s)
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	s2 := `new pass`
	check := bcrypt.CompareHashAndPassword(bs, []byte(s2))
	fmt.Println(check)

}
