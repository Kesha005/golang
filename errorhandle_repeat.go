package main


import (
	"fmt"
)


type Users struct{
	id int
	name string
	surname string
}



func FindOnMap(){

	var list[]string

	user := []Users{

		Users{1 , "Kerim" , "Saparow"},

		Users{2 , "Selbi" , "Saparowa"},

	}
	for _ , user  := range user{
		list = append(list, user.name)
	}
	fmt.Println(list)
}





func main(){
	FindOnMap()
}