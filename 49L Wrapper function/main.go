package main

import (
	"fmt"
	"log"
	"io/ioutil"
)

func main() {
	xb, err := readFile("49L Wrapper function\\output.txt")
	if err != nil {
		log.Fatalf("readFile in main: %s", err);
	}
	fmt.Println(xb);
	fmt.Println(string(xb));

}

func readFile(fn string) ([]byte, error) {
	xb, err := ioutil.ReadFile(fn);
	if err != nil {
		return nil, fmt.Errorf("error in readFile func: 5s", err);
	}
	return xb, nil;
}