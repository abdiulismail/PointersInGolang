package main

import "fmt"

//following a pointer means reading the value at the memory address that the pointer refers to
//its done using the asterik operator *

func main() {

	a := 100
	b := &a

	a++

	fmt.Println("first:", a)  //101
	fmt.Println("second", *b) //101  , dereferencing the pointer
}
