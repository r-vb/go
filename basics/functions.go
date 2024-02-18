package main

import (
	"fmt"
	"math"
)

func sayWelcome(n string) {
	fmt.Printf("Good Morning %v!\n", n)
}

func sayBye(n string) {
	fmt.Printf("Good Bye %v!\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleAre(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayWelcome("vivek")
	sayWelcome("Luigi")
	sayBye("mario")

	cycleNames([]string{"cloud", "raphel", "barrack"}, sayWelcome)
	cycleNames([]string{"cloud", "raphel", "barrack"}, sayBye)

	a1 := circleAre(10.5)
	a2 := circleAre(37)
	fmt.Println(a1, a2)
	fmt.Printf("Circle a1 area = %0.3f \nCircle a2 area = %0.2f\n", a1,a2)
}
