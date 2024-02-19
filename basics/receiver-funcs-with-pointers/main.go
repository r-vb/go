package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")

	mybill.addItem("tomato soup", 15.64)
	mybill.addItem("toffee pudding", 100.108)

	mybill.updateTip(10)

	fmt.Println(mybill.format())

}
