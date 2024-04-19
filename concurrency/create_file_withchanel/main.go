package main

import (
	"fmt"
	"log"
	"os"
	
)




// func Reader(ch chan string){

// }


func Writer(ch chan string){

	f, err := os.OpenFile(string(<-ch),os.O_WRONLY|os.O_CREATE, 0644)
	if err !=nil{
		log.Fatal("Error occured ", err)
		return 
	}
	defer f.Close()
	data := []byte("This is test data only")
	_,werr:= f.Write(data)
	if werr!=nil{
		log.Fatal(werr)
		return 
	}

	ch<-"file created successfully"



}

func main(){
	path:= make(chan string)
	go Writer(path)
	path<-"hi.txt"



	fmt.Println(<-path)

	
}