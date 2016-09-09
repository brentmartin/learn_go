package main

import (
	"fmt"
)

var a int
var b int

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

	fmt.Println("=================================================")
	fmt.Println("* Introducing pointers - var shares memory address that it is pointed too *")
	fmt.Println("")

	fmt.Println("     The setup:")
	a = 1
	b = 10
	fmt.Print("a's original value: ")
	fmt.Println(a)
	fmt.Print("b's original value: ")
	fmt.Println(b)
	fmt.Print("a's memory address: ")
	fmt.Println(&a)
	fmt.Print("b's memory address: ")
	fmt.Println(&b)
	fmt.Println("")
}
