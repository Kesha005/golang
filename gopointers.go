package main


import(
	"fmt"
)


func main(){

	variable := 34
	fmt.Println(variable)

	var  pointer *int
	
	pointer = &variable

	fmt.Println(pointer)

	*pointer=23
	fmt.Println(pointer)
	fmt.Println(variable)
}