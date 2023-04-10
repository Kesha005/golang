package main


import (
	"fmt"
)

func show(variable2 *int){
	*variable2 =2

}

func main(){
	variable1 := 1
	fmt.Println(variable1)
	variable2 := &variable1
	show(variable2)
	fmt.Println(variable1)
}