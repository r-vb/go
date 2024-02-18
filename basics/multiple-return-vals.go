package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1]) // [pos 0 but not 1 indicating first letter]
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn1, ln1 := getInitials("tiffa lockhart")
	fmt.Println(fn1, ln1)

	fn2, ln2 := getInitials("cloud striffe")
	fmt.Println(fn2, ln2)

	fn3, ln3 := getInitials("buffet")
	fmt.Println(fn3, ln3)
}
