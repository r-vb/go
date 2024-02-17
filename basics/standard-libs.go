package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Strings Package
	welcome := "nice to see you!"

	fmt.Println(strings.Contains(welcome, "to"))              // returns true or false value
	fmt.Println(strings.ReplaceAll(welcome, "nice", "happy")) // none of the functions change original string ; they create a new one with replacements made
	fmt.Println(strings.ToUpper(welcome))     // converts all letters in string to uppercase
	fmt.Println(strings.Index(welcome, "ee")) // returns first index where substring is found; -1 if not found
	fmt.Println(strings.Title("hello world")) // capitalizes first letter of each word
	fmt.Println(strings.Split(welcome, " "))

	// the original value is unchanged
	fmt.Println("original: ", welcome)

	//Sort Package
	ages := []int{45, 20, 85, 36, 29, 67, 55} //slice

	sort.Ints(ages)
	fmt.Println(ages) // prints the sorted slice

	index := sort.SearchInts(ages, 29)
	fmt.Println(index) // prints position where number is present in the sort slice

	names := []string{"yoshi", "rahul", "vivek", "mario", "vishwa"}
	
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "vishwa"))

}
