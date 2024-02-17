package main

import "fmt"

func main() {

	//strings
	var nameOne string = "Hello World!" //string literal is a sequence of characters enclosed in double quotes

	var nameTwo = "nice"

	var nameThree string
	nameThree = "to meet you."

	nameFour := "I am learning Go." //:= is used to declare and initialize a variable in the same statement. It's called a short declaration.

	fmt.Println(nameOne)
	fmt.Println("I am Alice", nameTwo) //%v prints the value
	fmt.Println(nameThree, nameFour)

	//ints
	var ageOne int = 20
	var ageTwo = 20
	ageThree := 20
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256

	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 25.69
	var scoreTwo float64 = -5243254355644531513545.351638464
	scoreThree := 1.5
	
	fmt.Println(scoreOne,scoreTwo,scoreThree)

}
