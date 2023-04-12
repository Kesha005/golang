package main


import(
	"fmt"
	"time"
)




func print_long(msg string){
	for i:=0 ; i < 3 ; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Its ",msg, " :" ,i)
	}
}


func main(){

	go print_long("Selbi") 

	print_long("Hello Kerim")

	



}