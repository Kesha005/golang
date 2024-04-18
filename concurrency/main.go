package main

import (
	"fmt"
	"sync"
	"time"
)


func Printer(thing string){
	for i:=0;  i<10; i++{
		fmt.Println(i, thing)
		time.Sleep(time.Second *1)
	}
}


func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	

	go func ()  {
		Printer("Hello ")
		wg.Done()
	}()

	wg.Wait()
	

}