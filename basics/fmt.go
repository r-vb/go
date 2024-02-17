package main

import "fmt"

func main() {

	//Print
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	//Println
	fmt.Println("hello Ninjas!")
	fmt.Println("goodbya ninjas!!")

	age := 35
	name := "shaun"
	fmt.Println("my age is", age, "and  my name is", name)

	//Printf (formatted strings) ----> %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name) //%v means print variable using default format
	fmt.Printf("my age is %q and my name is %q \n", age, name) // %q  uses double quotes around the string instead of single quotes.
	fmt.Printf("the type of age is %T \n", age)              //%T means print the type of the variable
	fmt.Printf("score %f points \n", 152.996)                //%f means float64
	fmt.Printf("score %0.2f points \n", 152.996)             //%.2f means float64 with two decimal places
	
	fmt.Printf("%s is %d years old\n", name, age) //%s is a placeholder for string and %d for integer

	//Sprintf (save formatted string)
	var strOne = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	strTwo := fmt.Sprintf("I am %s and I am %d years old", name, age)
	fmt.Println("The saved strings are:",strOne,strTwo)

}
