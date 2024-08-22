package main

import "fmt"

func main() {
	f();
	fmt.Println("Return nomally from f...");
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r);
		}
	}()

	fmt.Println("Calling g...");
	g(0);
	fmt.Println("Returned normally from g")
}

func g(i int) {
	if (i > 3) {
		fmt.Println("Panicking!!")
		panic(fmt.Sprintf("%v", i));
	}
	defer fmt.Println("Defer in g",i);
	fmt.Println("Printing in g",i);
	g(i+1);
}

/*
	# The main function calls f, which in turn calls g(0).
	# The function g prints a message and calls itself with an incremented value of i.
	# When i becomes greater than 3, g panics.
	# The panic causes the deferred functions in g to execute, printing messages in reverse order of the stack.
	# The panic is caught by the deferred anonymous function in f, which recovers from it and prints a recovery message.
	# The program then continues in f and returns to main, printing "Returned normally from f."
*/