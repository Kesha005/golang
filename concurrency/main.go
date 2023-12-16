package main

import (
	"fmt"
	"time"

	
)


func SayHello(msg string, cha chan string){

	time.Sleep(34)
	cha <-"Hi Kerim"
}



func main(){
				

	cha := make(chan string)

	go SayHello("hi", cha)

	mesg := <-cha

	fmt.Println(mesg)
}