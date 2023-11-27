package main 

import "fmt"



func main(){

	// defer fmt.Println("It's must be first!")
	// fmt.Println("Hello world")

	fmt.Println("Help something happening in program")


	panic("Ending program")

	fmt.Println("Program excecuted")


}