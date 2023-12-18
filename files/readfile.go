package main


import (
	"fmt"
	"os"
)


func Readfile(name string)(string,error){
	data, err := os.ReadFile(name)
	if err!=nil{
		return "",err
	}
	return string(data),nil
}


func main(){
	data, _:= Readfile("thirt.txt")
	fmt.Println(data)
}