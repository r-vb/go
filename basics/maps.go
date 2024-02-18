// maps datatype

package main

import (
	"fmt"
)

func main() {

	menu := map[string]float64{
		"soup":           4.99,
		"rice":           7.99,
		"salad":          9.66,
		"toffee pudding": 19.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "->", v)
	}

	// ints as key type
	phonebook := map[int]string{
		2765454151: "rahul",
		6396633514: "mom",
		2334654690: "dad",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[2765454151])

	phonebook[2765454151] = "bowser"
	fmt.Println(phonebook)

	phonebook[2334654690] = "yoshi"
	fmt.Println(phonebook)
}
