package main


import "fmt"


func update(msg *string){
	*msg = "Hello developer"
}


func main(){
	message := "Hello Kerim"
	fmt.Println(message)
	update(&message)

	fmt.Println(message)

	
}