package main

import("fmt")


type country interface{
	welcome(string)string
}

type english struct{}
type turkmen struct{}


func (r *english) welcome(name string)string{
	return fmt.Sprintf("Hello %s",name)
}

func (t *turkmen)welcome(name string)string{
	return fmt.Sprintf("Salam %s",name)
}

func privet(guest country , name string){
	fmt.Println(guest.welcome(name))
}

func main(){
	privet(&turkmen{},"Kerim")
	privet(&english{},"Kerim")
}