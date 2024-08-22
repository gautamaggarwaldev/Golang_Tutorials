package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string;
}

func (b book) String() string {
	return fmt.Sprint("Book title is ", b.title);
}

type count int;

func (c count) String() string {
	return fmt.Sprint("Number is ", strconv.Itoa(int(c)));
}

func Lofinfo(s fmt.Stringer) {
	log.Println("LOG from 125", s.String());
}

func main() {

	b := book {
		title : "7 khoon maaf",
	}

	var x count = 15;

	Lofinfo(b);
	Lofinfo(x);

}