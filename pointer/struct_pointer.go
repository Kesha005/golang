package main


import (
	"fmt"
)


type User struct{
	name string
	surname string
	age int
}

func (user *User)UpdateUser(data User){
	user.name = data.name
	user.surname =data.surname
	user.age= data.age
}

func main(){
	fmt.Println("Hello User how are you user")


	user := User{"Kerim","Saparow",18}

	fmt.Println("Username: ",user.name)

	user.UpdateUser(User{"Kakamyrat","Saparow",16})

	fmt.Println(user.name)

}