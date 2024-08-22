package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"fmt"
)

type person struct {
	first string;
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {

	p := person {
		first : "garima",
	}


	f, err := os.Create("42eWriting into file either or a buffer\\output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer;

	p.writeOut(f);
	p.writeOut(&b);
	fmt.Println(b.String());
}