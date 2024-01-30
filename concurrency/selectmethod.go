package main  


import (
	"fmt"
	"time"
)




func getSurname(ch chan string){
	time.Sleep(2 *time.Second)
	ch<-"Saparow"
}


func getName(ch chan string){
	time.Sleep(2 *time.Second)
	ch<-"Kerim"
}


func main(){
	first := make(chan string)
	second := make(chan  string)
	go getName(first)
	go getSurname(second)
	select{
	case firstdata := <-first:
		fmt.Println(firstdata)
	case seconddata:= <-second:
		fmt.Println(seconddata)
	}
}

