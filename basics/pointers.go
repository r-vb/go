package main

import (
	"fmt"
)

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"

	// updateName(name)

	fmt.Println("Memory address of the 'name' is: ", &name)

	m := &name
	fmt.Println("Memory address: ", m)
	fmt.Println("Value at the above memory address: ", *m)

	fmt.Println(name) // tifa

	updateName(m)

	fmt.Println(name)  // wedge
}

/*
|--name--|---m----|
|  0x001 | 0x002  |
|--------|--------|
| "tifa" | p0x001 |
|--------|--------|
*/
