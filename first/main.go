package main

import (
	"errors"
	"fmt"
)


func divide(x1, x2 int)(int,error){

	if x2 ==0{
		return 0, errors.New("Zero division error ")
	}
	result := x1/x2

	return result,nil
	
}

func main(){
	fmt.Println("Hello user enter two numbers to divide")
	fmt.Println("Enter first: ")
	var first, second int
	fmt.Scanln(&first)
	fmt.Println("Enter second one: ")
	fmt.Scanln(&second)
	result,err:= divide(first, second)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
}
