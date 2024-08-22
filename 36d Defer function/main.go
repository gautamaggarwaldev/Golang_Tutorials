package main

import "fmt"

/*
A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.
*/

func foo() {
	fmt.Println("hey, I am foo.")
}

func bar() {
	fmt.Println("hey, I am bar.")
}

func main() {
	defer foo();  //2
	bar();  //1
}