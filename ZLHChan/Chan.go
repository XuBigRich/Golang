package main

import "fmt"

func main() {
	ch := make(chan int32, 1)
	go say(ch)
	ch <- 1
	ch <- 2
}
func say(ch chan int32) {
	fmt.Print(<-ch)
	fmt.Print(<-ch)
}
