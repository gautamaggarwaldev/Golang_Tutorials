package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{8, 7, 6, 14, 5, 1, 96, 88, 45, 3, 23, 2, 6, 9, 5}
	fmt.Println(a);
	sort.Ints(a);
	fmt.Println(a);

	b := []string{"vkf", "asdfg", "werty", "lkj", "poiuh", "cf"};
	fmt.Println(b);
	sort.Strings(b);
	fmt.Println(b);
}