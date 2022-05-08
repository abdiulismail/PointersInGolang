package main

import "fmt"

//given that pointers store memory location , it is possible to create a pointer whose value is the memory address of another pointer

func main() {

	first := 100
	second := &first
	third := &second

	fmt.Println(first)
	fmt.Println(*second)
	fmt.Println(**third)
}
