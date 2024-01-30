package main

import (
	
	"fmt"
)


func returnName(ch chan string){
	ch<-"Kesha"
	ch<-"Saparow"
	ch<-"Are you ok"
}


func main(){
	ch := make(chan string, 2)
	go returnName(ch)
	
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)


	ch1 := make(chan int,2)
	ch1<-2
	ch1<-1
	
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	
}