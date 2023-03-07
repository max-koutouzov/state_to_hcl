package main

import "fmt"

func main() {
	var p = [3]int{4, 7, 23}
	fmt.Println(p)

	f := [...]int{4, 6, 7, 23, 34, 53, 24}
	fmt.Println(f)

	for i := 0; i <= len(f)-1; i++ {
		fmt.Println(f[i])
	}

	n := make([]int, 3)
	n[0] = 1
	fmt.Println(len(n))
	n = append(n, 4)
	fmt.Println(len(n))
}
