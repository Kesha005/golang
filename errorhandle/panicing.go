package main

import (
	"fmt"

)


func HandlePanic(){
	a := recover()

	if a !=nil{
		fmt.Println("Recover ", a)
	}

}


func div(x, y int)float64{

	defer HandlePanic()
	if y==0{
		panic("Division to zero is error")

	}else{
		num:=  x/y
		return float64(num)
	}
}



func main(){
	fmt.Println(div(2,4))
	fmt.Println(div(2,2))
	fmt.Println(div(2,0))
}