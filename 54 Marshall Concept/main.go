package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string;
	Lname string;
	Age int;
}

func main() {
	p1 := person{
		Fname: "garima",
		Lname : "govil",
		Age: 27,
	}

	p2 := person{
		Fname: "Jenny",
		Lname : "Watson",
		Age: 17,
	}

	people := []person{p1,p2};

	fmt.Println(people);

/*
	func Marshal :-

	func Marshal(v any) ([]byte, error)

	Marshal returns the JSON encoding of v.

	Marshal traverses the value v recursively. If an encountered value implements Marshaler and is not a nil pointer, Marshal calls [Marshaler.MarshalJSON] to produce JSON. If no [Marshaler.MarshalJSON] method is present but the value implements encoding.TextMarshaler instead, Marshal calls encoding.TextMarshaler.MarshalText and encodes the result as a JSON string. The nil pointer exception is not strictly necessary but mimics a similar, necessary exception in the behavior of [Unmarshaler.UnmarshalJSON].
	
*/

	bs, err := json.Marshal(people); // json.Marshal()
	if err != nil {
		fmt.Println(err);
	}
	fmt.Println(string(bs));
}