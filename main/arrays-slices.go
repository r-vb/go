package main

import "fmt"

func main() {

	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}
	fmt.Println(ages, len(ages))

	names := [4]string{"mario", "rahul", "vishwa", "vivek"}
	names[0] = "yuigi"
	fmt.Println(names, len(names))

	// slices (they use array under the hood)
	var scores = []int{100, 50, 60}
	scores[1] = 30
	scores = append(scores, 70)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3] // inclusive of 1 but not 3
	rangeTwo := names[1:]  // inclusive of 1 to end to array
	rangeThree := names[:3] // start till .. 3 but exclusive of 3
	fmt.Println(rangeOne,rangeTwo,rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println("After appending koopa : ", rangeOne)

}
