package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{name: "Bob", age: 22})
	fmt.Println(person{name: "Charlene", age: 24})
}
