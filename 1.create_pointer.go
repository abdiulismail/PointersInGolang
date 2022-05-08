package main

import "fmt"

func main() {

	Grade1 := 20
	Grade2 := &Grade1 //Grade2 is assigned the memory
	//address of Grade1

	Grade1++ //This does not affect Grade2 since go does not support
	//pointer arithmetic

	fmt.Println(Grade1)
	fmt.Println(*Grade2)

}
