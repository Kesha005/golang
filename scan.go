package main

import("fmt")



func main(){
	
	var user_data int
	fmt.Scan(&user_data, "Enter your name ")
	if(user_data < 9 ){
		fmt.Println("its so small")
	}
	fmt.Println(user_data)
}