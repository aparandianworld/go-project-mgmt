package main

// Default variables here
var (
	itemPrice    = 100
	itemDiscount = 10
)

func BlackFridayDiscount() int {
	return itemDiscount * 2
}
