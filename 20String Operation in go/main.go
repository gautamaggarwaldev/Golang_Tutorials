package main

import "fmt"

func main() {
	str := "Hello World!"
	fmt.Println(len(str))
	fmt.Println(str[0], str[11])  // it print the ascii value of the alphabet which is present in the index 0 and index 11
	fmt.Println(str[0:9]) //it print the character from 0 to 9 index

	fmt.Println(str[4:]) //start print from 4th index upto last
	fmt.Println(str[:7]) //start print from 0th index upto 7th
	fmt.Println(str[:]) // print complete string
	fmt.Println("Gautam " + str[0:5])

}