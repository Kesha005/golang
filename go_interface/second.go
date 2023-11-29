package main


import "fmt"


type User struct{
	id int
	name string
	lname string
	role int // 1 admin 0 superadmin
}

type Student struct{
	id int 
	name, lname string
	course int
	faculty string
}


type Auth  interface{
	Login()
}

func (u User)Login(){
	if u.role ==1{
		fmt.Println("Welcome ", u.name," you are site superadmin")
	}else{
		fmt.Println("Welcome ", u.name," you a admin")
	}
}


func (s Student) Login(){
	fmt.Println("Hi ", s.name," you a ",s.course," sudent of ",s.faculty," faculty")
}


func  Authorize(a Auth){
	a.Login()
}


func main(){
	user := User{1,"Kerim", "Saparow",1}
	user1 := User{1,"Kakamyrat", "Saparow",0}
	student := Student{1,"Selbi","Saparowa",2,"itp"}
	user.Login()
	user1.Login()
	student.Login()
}