package main

import "fmt"

func main() {
	new_channel := make(chan string)
	go func() { new_channel <- "ping" }()
	msg := <-new_channel
	fmt.Println(msg)
}
