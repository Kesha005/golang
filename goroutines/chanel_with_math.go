package main

import (
	"fmt"
	"math"
)




func main(){
	chanel := make(chan float64)
	go power(3,4,chanel)

	fmt.Print("The result is ", <-chanel)
}



func power(x float64 ,y float64,chanel chan float64){
	result := math.Pow(x,y)
	chanel <- result
}

