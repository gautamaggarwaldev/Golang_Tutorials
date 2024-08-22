package main

import (
	"fmt"
	"unicode/utf8"
)
func main() {
	s := "hello world!"
	fmt.Println(len(s))

	for i:=0; i<len(s); i++ {  //This loop generates the hex values of all the bytes that constitute the code points in s.
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count : ", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)  //print unicode and index of the characters
    }
}