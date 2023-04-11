package main


import("fmt")


type User interface{
	show()
}


func show(user User){
	fmt.Println(user)
}


func main(){
	var user User
	user.show()
}