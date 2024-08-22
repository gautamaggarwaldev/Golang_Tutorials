package main

import "fmt"

type person struct {
	first string;
}

type secretAgent struct {
	person;
	ltk bool;
}

func (p person) speak() {
	fmt.Println("I am", p.first);
}

func (sa secretAgent) speak() {
	fmt.Println("I am", sa.first);
}

type human interface {
	speak();
}

func say(h human) {
	h.speak();
}

func main() {
	sa1 := secretAgent {
		person: person {
			first: "james",
		},
		ltk: true,
	}

	p2 := person {
		first: "garima",
	}

	say(sa1);
	say(p2);
}