/*
	var channel_name chan type

	channel_name := make(chan type)
*/

package main

import "fmt"
func main() {
	/*declaration of channel*/

	var mychannel chan int;
	//mychannel := make(chan int);
	fmt.Println("value of channel:",mychannel);
	fmt.Printf("type of channel: %T",mychannel);


	/*
		chan1 <- value  //send
		var := <-chan1
		<- chan1
	*/

	channel := make(chan int);
	fmt.Println("\nHello from main");
	go run(channel);
	channel <- 12;
	fmt.Println("end")


	/*close channel*/
	ch := make(chan int);
	close(ch);
	ele,ok := <- ch;
	fmt.Println(ele,ok);
}

func run(channel chan int) {
	fmt.Println(100 * <-channel);
}