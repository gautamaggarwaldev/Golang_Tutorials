package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("Hello ");
	fmt.Println(b.String());
	b.WriteString("Duniya!");
	fmt.Println(b.String());
	b.Reset();
	b.WriteString("my name is khan.");
	b.Write([]byte("happy happy"));
	fmt.Println(b.String());
}