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

	fmt.Println("     Now lets point b to a ")
	fmt.Println("> Error: '*b = &a' won't work, we'll have to recreate it from scratch and point it to a with 'var b *int = &a'")
	fmt.Println("")

	var b *int = &a
	fmt.Println("     after recreating b and pointing it to a in the process:")
	fmt.Print("a's original value: ")
	fmt.Println(a)
	fmt.Print("b's recreated value: ")
	fmt.Println(b)
	fmt.Print("a's original memory address: ")
	fmt.Println(&a)
	fmt.Print("b's recreated memory address: ")
	fmt.Println(&b)
	fmt.Println("> Interesting, it seems that b's value is just a's address")
	fmt.Println("")

	fmt.Println("     lets try doing something like printing 'pointer b' vs just 'b' ('*b' vs 'b')")
	fmt.Print("b's value (with pointer): ")
	fmt.Println(*b)

	fmt.Println("")
	fmt.Println("notes: pointers must be assigned when initialized, can't be reassigned as a pointer later when was previously not a pointer")

	fmt.Println("")
	fmt.Println("=================================================")
	fmt.Println("* playing with the values of the address being pointed to *")
	fmt.Println("")

	fmt.Println("     The setup: using b to point to a")
	a = 2
	var c *int = &a
	fmt.Print("a's original value: ")
	fmt.Println(a)
	fmt.Print("a's memory address: ")
	fmt.Println(&a)
	fmt.Print("pointer c's value ")
	fmt.Println(*c)
	fmt.Println("")
}
