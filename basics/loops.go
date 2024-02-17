package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of x is:", i)
	}

	// looping through slice of strings or slice of numbers
	names := []string{"alice", "mario", "vishwa", "boston", "rahul"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("the value of index %v is %v \n", index, value)
	}

	for _, value := range names { // if we don't want to use index
		fmt.Printf("the value is %v \n", value) // ll'rly if we don't want to use value replace with _
		value = "new string"
	}	// altering value here does not update the original slice as it creates a new local copy of the variable in each iteration
	fmt.Println(names)	
}
