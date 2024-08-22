package main

import "fmt"

// func (r receiver) identifier (p parameter(s)) (return(s)) { ... }

func foo() {
	fmt.Println("hey, I am foo");
}

func name(s string) {
	fmt.Println("My name is", s);
}

func agent(s string) string {
	return fmt.Sprint("her name is ",s);
}

func student(n string, age int) (string, int) {
	age += 50;
	return fmt.Sprint("name : ",n) , age;
}

func main() {
	foo();
	name("garima");
	s := agent("Govil ji");
	fmt.Println(s);
	a, b := student("gautam", 21);
	fmt.Println(a,b);
}