package main

import (
	"fmt"
)

func main() {
	totalDiscount := Discount(itemPrice, itemDiscount)
	blackFridayDiscount := BlackFridayDiscount()

	finalPrice := itemPrice - int(totalDiscount)
	blackFridayFinalPrice := itemPrice - int(blackFridayDiscount)

	fmt.Println("Item Price:", itemPrice)
	fmt.Println("Item Discount:", itemDiscount)
	fmt.Println("Total Discount:", totalDiscount)
	fmt.Println("Final Price:", finalPrice)
	fmt.Println("Black Friday Discount:", blackFridayDiscount)
	fmt.Println("Black Friday Final Price:", blackFridayFinalPrice)
}
