package main

import "fmt"

func isDelta(y *int) {
	*y = 50;
}

func sliceDelta(s1 []int) {
	s1[1] = 1000;
}

func mapDelta(m1 map[string]int, s string) {
	m1[s] = 400;
}

func main() {

	x := 42
	fmt.Println(x);
	isDelta(&x);
	fmt.Println(x);

	slice := []int{1,2,3,4,5,6};
	fmt.Println(slice);
	sliceDelta(slice);
	fmt.Println(slice);

	m := make(map[string]int);
	m["james"] = 200;
	fmt.Println(m);
	mapDelta(m,"james");
	fmt.Println(m);


	


}