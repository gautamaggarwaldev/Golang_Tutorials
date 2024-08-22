package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func main() {

	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err);
	}
	fmt.Println(s);
	fmt.Println(bs);

	loginPword1 := `password1234`;

	err = bcrypt.ComparehashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("You can't login");
		return;
	}

	fmt.Println("You are logged in");

}