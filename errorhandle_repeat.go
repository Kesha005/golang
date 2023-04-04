package main

import (
	"errors"
	"fmt"
)


type Users struct{
	id int
	name string
	surname string
}


func FindOnMap(id int)(error ,string){



	users := []Users{

		Users{1 , "Kerim" , "Saparow"},

		Users{2 , "Selbi" , "Saparowa"},

	}
	var user_name string
	for _ , user  := range users{
		if(user.id==id){
			user_name = user.name
		}

		return errors.New("This user not found"), "not found"
		
	}
	return nil,user_name

}


func main(){

	err , user_name := FindOnMap(5)
	if(err !=nil){
		fmt.Println(err)
	}
	fmt.Println(user_name)
}