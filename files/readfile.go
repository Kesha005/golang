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

func ControlIsThere()(string){
	res,err := os.UserHomeDir()
	if err!=nil{
		fmt.Println(err.Error())
	}
	return res
}



func main(){
	data, _:= Readfile("thirt.txt")
	fmt.Println(data)
	fmt.Println(ControlIsThere())
}