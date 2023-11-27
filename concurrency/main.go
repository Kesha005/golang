package main


import "fmt"




func sayhello(){
	fmt.Println("Hello Kerim")
}


func main(){


	go sayhello()
	fmt.Println("Hello to concurrency")
}
