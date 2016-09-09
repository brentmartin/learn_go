package main

import (
	"fmt"
)

var a int
func main() {
	a = 1
	fmt.Println("* Memory address does not change when reassigned *")
	fmt.Println("")

	fmt.Println("     The setup:")
	fmt.Print("original address: ")
	fmt.Println(&a)
	fmt.Print("original value: ")
	fmt.Println(a)

	a = 2
	fmt.Println("     Now lets reassign value from 1 to 2 - value changes, address stays")
	fmt.Print("post-reassigned address: ")
	fmt.Println(&a)
	fmt.Print("post-reassigned value: ")
	fmt.Println(a)

}
