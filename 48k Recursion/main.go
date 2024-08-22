package main

import "fmt"

func main() {

	fmt.Println(factorial(20));
}

func factorial(n int64) int64 {
	if(n==0 || n==1) {
		return 1;
	}
	return n * factorial(n-1)
}