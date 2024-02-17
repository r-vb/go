package main

import (
	"fmt"
)

func main() {
	age := 45

	//comparision operators
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	// simple if, else if, else
	if age < 30 {
		fmt.Println("Age is less than 30!")
	} else if age > 30 {
		fmt.Println(age, "is greater than", 30)
	} else {
		fmt.Println("Exactly 30.")
	}

	// continue & break usage
	names := []string{"alice", "mario", "vishwa", "boston", "rahul"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v\n",index,value)
	}
}
