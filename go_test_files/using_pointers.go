package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x
	fmt.Printf("Printing pointer of p: %v\n", *p)
	*p = 2
	fmt.Printf("Assigning 2 to p pointer: %v\n", x)
}
