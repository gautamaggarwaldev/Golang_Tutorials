package main

import "fmt"

type person struct {
	first string;
	last string;
	age int;
}

func main() {
	p1 := struct {
		first string;
		last string;
		age int;
	} {
		first: "garima",
		last: "govil",
		age: 27,
	}

	p2 := person {
		first: "gautam",
		last: "aggarwal",
		age: 20,
	}

	fmt.Println(p1);
	fmt.Println(p2);
  
	fmt.Printf("%T\n", p1); // struct { first string; last string; age int }
	fmt.Printf("%T\n", p2);
/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
	fmt.Printf("%T\n",p2);
	fmt.Printf("%#v\n",p2);

	p2 = p1;

	fmt.Printf("%T\n",p2);
	fmt.Printf("%#v\n",p2);
}