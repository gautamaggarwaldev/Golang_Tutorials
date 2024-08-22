package main

import "fmt"

func addOne(v int) int { //pass by value ---> value semantic
	return v + 1;
}

func addOne1(v *int)  { //pass by reference ---> pointer semantic
	*v += 1;
}

func main() {
	x := 1
	fmt.Println(x); //1
	fmt.Println(addOne(x));//2
	fmt.Println(x);//1

	y := 1
	
	fmt.Println(y);//1
	addOne1(&y); 
	fmt.Println(y);//2
}