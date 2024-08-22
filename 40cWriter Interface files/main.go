package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("40cWriter Interface files\\output.txt");
	if err != nil {
		log.Fatalf("error %s", err);
	}
	defer f.Close();

	s := []byte("Hello Duniya!!");

	_,err = f.Write(s);
	if err != nil {
		log.Fatalf("error %s", err);
	}
}



/*
	# Converting a string to a byte slice :-
	str := "hello"
	bytes := []byte(str)


	# Converting a byte slice to a string :-
	bytes := []byte {72,101,108,108,111}
	str := string(bytes)
*/