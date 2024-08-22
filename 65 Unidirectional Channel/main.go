package main
import "fmt"
func main() {
	ch1 := make(chan string);
	go convert(ch1);
	fmt.Println(<-ch1);
}

func convert(ch chan<- string) { //send only channel we can't receive the value from it
	ch <- "garima";
	// <- ch; throws error
}