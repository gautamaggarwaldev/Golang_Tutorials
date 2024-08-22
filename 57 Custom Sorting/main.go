package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string;
	age int;
}

type ByAge []person
type ByName []person

func(a ByAge) Len() int {
	return len(a)
}
func(a ByAge) Swap(i,j int) {
	a[i], a[j] = a[j], a[i];
}
func(a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age;
}



func(bn ByName) Len() int {
	return len(bn)
}
func(bn ByName) Swap(i,j int) {
	bn[i], bn[j] = bn[j], bn[i];
}
func(bn ByName) Less(i, j int) bool {
	return bn[i].name < bn[j].name;
}


func main() {
	p1 := person{name:"A", age:24};
	p2 := person{name:"D", age:23};
	p3 := person{name:"R", age:25};
	p4 := person{name:"Q", age:22};
	p5 := person{name:"H", age:21};

	people := []person{p1,p2,p3,p4,p5};

	fmt.Println(people);
	fmt.Println("Sort BY AGE :-");
	sort.Sort(ByAge(people));
	fmt.Println(people);
	
	fmt.Println("Sort BY NAME :-");
	sort.Sort(ByName(people));
	fmt.Println(people);


}