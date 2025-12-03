package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed numbers.txt
var data []byte

func main() {
	lines := strings.Split(string(data), "\n")
	var sum int

	for _, line := range lines {
		if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("%s is not a valid number\n", line)
				continue
			}
			sum += num

		}
	}

	fmt.Println("Sum of lines (numbers):", sum)
}
