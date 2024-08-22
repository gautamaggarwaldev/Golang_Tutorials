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
	y := []int{1,2,3,4,5,6,7,8,9,10}; //slice

	x :=  sum(y...); //unfluring a slice

	fmt.Printf("sum : %v\n",x);
}