package defaults

import (
	"fmt"
	"math/rand"
)

// Default variables here
var (
	ItemPrice    = 100
	ItemDiscount = 10
)

func BlackFridayDiscount() int {
	return ItemDiscount * 2
}

func randomDiscount() int {
	min := 1
	max := 50
	ItemDiscount = rand.Intn(max-min+1) + min
	return ItemDiscount
}

func RandomDiscount() int {
	return randomDiscount()
}

func init() {
	fmt.Println("Init() in defaults.go called number 1...")
}

func init() {
	fmt.Println("Init() in defaults.go called number 2...")
}
