package main

import "fmt"

var points = []int{20, 90, 100, 45, 70} // points slice

func sayHello(n string) {
	fmt.Println("hello", n)
}

// func showScore(score float64) {
// 	fmt.Println("you scored these points: ", score)
// }

func showScore() {
	fmt.Println("you scored these points: ", score)
}