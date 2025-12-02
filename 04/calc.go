package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func Discount(a, b int) float64 {
	return float64(a*b) / 100
}

func init() {
	fmt.Println("Init() in calc.go called...")
}
