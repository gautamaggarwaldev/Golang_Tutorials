package main

import "fmt"

func main() {
	ch := make(chan string)

	go (func() {
		ch <- "value 1"
		ch <- "value 2"
		close(ch)
	})()

	for i := range ch {
		fmt.Println(i);
	}

}