package main
import "fmt"

type dog struct {
	first string;
}

func (d dog) walk() {
	fmt.Println("His Name is", d.first , "and he is walking.");
}

func (d *dog) run() {
	d.first = "taruna";
	fmt.Println("Her Name is", d.first , "and she is running.");
}

type youngin interface {
	walk();
	run();
}

func youngRun(y youngin){
	y.run();
}

func main() {
	d1 := dog{"pradeep"}
	d1.walk();
	d1.run();
	// youngRun(d1);----> throws error

	d2 := &dog{"clg"};
	d2.walk();
	d2.run();
	youngRun(d2);
}