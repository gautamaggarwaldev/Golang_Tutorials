package main

import "fmt"

func main() {
	ch := make(chan string);
	go fun(ch);
	for {
		resp, ok := <- ch
		if(ok == false) {
			fmt.Println("channel close", ok);
			break;
		} 
		fmt.Println("channel open", resp,ok);
	}


	ch2 := make(chan string, 6);
	ch2 <- "A";
	ch2 <- "B";
	ch2 <- "C";
	ch2 <- "D";
	ch2 <- "E";
	ch2 <- "F";
	fmt.Println("length of channel ch2 is", cap(ch2));
}

func fun(ch chan string) {
	for i := 0; i < 3; i++ {
		ch <- "garima";
	}
	close(ch);
}