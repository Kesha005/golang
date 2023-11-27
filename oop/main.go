package main



import(
	"fmt"
)



type User struct{
	name, surname string
	age int
}



func (u User)GetName()string{
	return u.name
}


func (u User)GetFullName()string{
	return u.name+" "+u.surname
}



func main(){
	user := User{"Kerimberdi","Saparow",18}
	fmt.Println(user.GetFullName())
	
}