package main

import "fmt"

func main() {
	// Bidirectional channel
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sending(ch1);
	valueFromch1 := <-ch1;
	fmt.Println(valueFromch1);
	go receiving(ch2);
	ch2 <- valueFromch1;
}

func sending(ch chan string) {
	ch <- "garima"
}
func receiving(ch chan string) {
	fmt.Println(<-ch);
}