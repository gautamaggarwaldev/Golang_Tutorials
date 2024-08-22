package main

import (
	"fmt"
	"time"
)

func main() {
	fun();
	go fun();
}

func fun() {
	for i := 0; i < 10; i++ {
		time.Sleep(1*time.Second)
		fmt.Println(i);
	}
}