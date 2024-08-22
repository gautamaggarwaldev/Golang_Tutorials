package main

import "fmt"

func main() {
	foo();

	//func(p parameter(s)) (r return(s)) {code}

	x := func() {
		fmt.Println("Anonymous Function run...");
	}

	y := func(s string){
		fmt.Println("From anonymous func my name is ",s);
	}

	x();
	y("garima");
}

func foo() {
	fmt.Println("foo ran");
}
