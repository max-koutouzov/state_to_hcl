package main

import "fmt"

func main() {
	for c := 0; c <= 30; c++ {
		fmt.Println(c)
		switch c {
		case 2:
			fmt.Println("fizz")

		}
	}
}
