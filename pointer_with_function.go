package main 


import(
	"fmt"
)


type User struct {
	name string
	email string
	year int 
}



func (user *User)update_user(name string , email string , year int){
	user.name = name
	user.email = email
	user.year = year
	
}



func (user User)show_user(){
	fmt.Println("user name is " , user.name)
	fmt.Println("User email is " ,user.email)
	fmt.Println("User ",user.year , " old")
}



func main(){
	user := User{"Kerim","kerim@gmail.com",17}
	user.show_user()
	user.update_user("Myrat","myrat@gmail.com",19)
	fmt.Println(user)
}

