package main


import "fmt"


func update(msg *string){
	*msg = "Hello developer"
}


func delete(msg *string){
	*msg = ""
}


func main(){
	message := "Hello Kerim"
	fmt.Println(message)
	update(&message)
	delete(&message)
	fmt.Println(message)
}