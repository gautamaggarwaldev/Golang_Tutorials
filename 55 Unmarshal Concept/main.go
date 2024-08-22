package main

import (
	"encoding/json"
	"fmt"
)
	
type person struct {
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"Fname":"garima","Lname":"govil","Age":27},{"Fname":"Jenny","Lname":"Watson","Age":17}]`;
	bs := []byte(s);

	fmt.Printf("%T\n",s);
	fmt.Printf("%T\n",bs);

	var people []person;

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err);
	}

	fmt.Println("all of the data", people);

	for i, v := range people {
		fmt.Println("\nPerson Number", i);
		fmt.Println(v.Fname, v.Lname, v.Age);
	}
}