package main

import "fmt"

type person struct {
	first string;
	last string;
	age int;
}

type secretAgent struct {
	person;
	ltk bool;
	first string;
}

func main() {

	sA := secretAgent {
		person: person{
			first: "garima",
			last: "govil",
			age: 27,
		},
		first : "gg",
		ltk: true,
	}

	p2 := person {
		first: "gautam",
		last: "aggarwal",
		age: 20,
	}

	fmt.Println(sA);
	fmt.Println(p2);

	fmt.Printf("%T\t %#v\n",sA,sA);

	fmt.Println(sA.first, sA.last, sA.age, sA.person, sA.ltk);

	fmt.Println(sA.first, sA.person.first); // gg garima
}