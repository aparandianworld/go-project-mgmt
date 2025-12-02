package main

import (
	"fmt"

	"github.com/aparandianworld/go-project-mgmt/03/defaults"
)

func main() {
	totalDiscount := defaults.Discount(defaults.ItemPrice, defaults.ItemDiscount)
	blackFridayDiscount := defaults.BlackFridayDiscount()
	randomDiscount := defaults.RandomDiscount()

	finalPrice := defaults.ItemPrice - int(totalDiscount)
	blackFridayFinalPrice := defaults.ItemPrice - int(blackFridayDiscount)
	randomFinalPrice := defaults.ItemPrice - int(randomDiscount)

	fmt.Println("Item Price:", defaults.ItemPrice)
	fmt.Println("Item Discount:", defaults.ItemDiscount)
	fmt.Println("Total Discount:", totalDiscount)
	fmt.Println("Final Price:", finalPrice)
	fmt.Println("Black Friday Discount:", blackFridayDiscount)
	fmt.Println("Black Friday Final Price:", blackFridayFinalPrice)
	fmt.Println("Random Discount:", randomDiscount)
	fmt.Println("Random Final Price:", randomFinalPrice)
}
