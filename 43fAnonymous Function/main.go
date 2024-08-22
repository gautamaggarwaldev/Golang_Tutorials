package main

import "fmt"

func main() {
	foo();

	//func(p parameter(s)) (r return(s)) {code}

	func() {
		fmt.Println("Anonymous Function run...");
	}()

	func(s string){
		fmt.Println("From anonymous func my name is ",s);
	}("Garima")
}

func foo() {
	fmt.Println("foo ran");
}
