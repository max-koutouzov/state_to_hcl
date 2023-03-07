package main

import "fmt"

func boolConversion(value bool) int {
	if value {
		return 1
	}
	return 0
}

func main() {
	fmt.Printf("True bool statement: %b\n", boolConversion(true))
	fmt.Printf("False bool statement %b\n", boolConversion(false))
}
