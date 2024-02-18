/* Go is a Pass-by-value Language.<br>
// When you pass a variable to a function in Go, the value of that variable is copied.
__________________________
|__Group A__|__Group B___|
|*. Strings |*. Slices	 |
|*. Ints	|*. Maps	 |
|*. Floats  |*. Functions|
|*. Booleans|			 |
|*. Arrays  |			 |
|*. Structs |			 |
|------------------------|
*/

package main

import (
	"fmt"
)

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["pizza"] = 15.09
}

func main() {
	// group A types --> strings, ints, bools, floats, arrays, structs
	name := "tiffa"

	name = updateName(name) // name is passed by value so it's not updated outside this func
	fmt.Println(name)

	// group B types --> slices, maps, functions
	menu := map[string]float64{
		"ice cream": 30.5,
		"pie":       5.95,
	}

	updateMenu(menu)
	fmt.Println(menu)
}

/*
___________________________________________________
|__Non-Pointer Values__|__Pointer Wrapper Values__|
|	*. Strings 		   |	*. Slices	 		  |
|	*. Ints			   |	*. Maps	 			  |
|	*. Floats  		   |	*. Functions		  |
|	*. Booleans		   |			 			  |
|	*. Arrays  		   |						  |
|	*. Structs 		   |						  |
|----------------------|--------------------------|

*/
