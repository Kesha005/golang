package main

import("fmt")


type User struct{
	name string
	email string
	address string
}


type UserInterface interface{

	delete_user()
	show_user()
	update_user()
}


func (user *User)show_user(){
	fmt.Println(user.name)
}


func (user *User) update_user(name string , email string , address string){
	user.email = email
	user.name = name
	user.address = address
}

func delete_user(){

}


func main(){
	user := &User{"Kerim","kerim@gmail.com","Mekan28"}
	user.show_user()
	user.update_user("Myrat ", "myrat@gmail.com","Ashgabat 28")
	fmt.Println(user)
	user.show_user()
}