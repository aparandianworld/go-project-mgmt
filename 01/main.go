package main

import (
	"fmt"
)

func main() {
	itemPrice := 100
	itemDiscount := 10

	totalDiscount := Discount(itemPrice, itemDiscount)
	finalPrice := itemPrice - int(totalDiscount)

	fmt.Println("Item Price:", itemPrice)
	fmt.Println("Item Discount:", itemDiscount)
	fmt.Println("Total Discount:", totalDiscount)
	fmt.Println("Final Price:", finalPrice)
}
