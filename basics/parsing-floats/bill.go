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
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
func (b *bill) format() string {
	fs := "\nBill Breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...₹%v \n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...₹%0.2f \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...₹%0.2f \n", "total:", total+b.tip)

	return fs
}

// update the tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip // or (*b).tip = tip --> de-referencing
}

// add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}