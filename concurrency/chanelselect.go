package main

import (
	"fmt"
	"time"
	
)


func main(){
	numeral := make(chan int)
	Stringval := make(chan string)
	go GetNumral(numeral)
	go getString(Stringval)

	select{
	case firstChanel := <-numeral:
		fmt.Println(firstChanel)
	
	case secondChanel := <-Stringval:
		fmt.Println(secondChanel)
	}

}



func GetNumral(nm chan int){
	time.Sleep(2 * time.Second)
	nm <-12
}


func getString(st chan string){
	time.Sleep(2 * time.Second)
	st <-"Hello world"
}