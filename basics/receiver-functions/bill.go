package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// amke new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"lafda": 21.56, "choco": 10.96, "lotte": 30.84},
		tip:   0,
	}

	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-10v ...₹%v \n", k+":", v)
		total += v
	}

	// total
	fs += fmt.Sprintf("%-10v ...₹%0.2f", "total:", total)

	return fs
}
