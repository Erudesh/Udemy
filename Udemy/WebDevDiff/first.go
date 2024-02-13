package main

import "fmt"

type names struct {
	fname string
	sname string
	tname string
}

func (n names) speak() {
	fmt.Println(n.fname, n.sname, n.tname, "glad to see you!")
}

func main() {
	n7 := names{
		"Anthony",
		"Degrues",
		"Scott",
	}
	n7.speak()
}
