// Variadic parameter = when you create a func whic takes an unlimited number of arguments, is known as a VP.

//we can call the variadic function with passing zero or more arguments. 

package main

import "fmt"

func sum(i ...int) int {
	fmt.Println(i);
	fmt.Printf("%T\n", i); // []int ---> slice

	n := 0;
	for _,v := range i {
		n += v;
	}
	return n;
}

func main() {
	x :=  sum(1,2,3,4,5,6,7,8,9,10);
	fmt.Printf("sum : %v\n",x);
}