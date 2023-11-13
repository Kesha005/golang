package main

import "fmt"


func main(){
	// This is an array example
	// first_array := [...]int{1,2,4,5,6,7,8}
	// fmt.Println(first_array)
	// for i:=0; i<len(first_array); i++{
	// 	fmt.Println(first_array[i])
	// }

	first_map := make(map[int]string)

	first_map[0] = "Kerim"
	first_map[1] = "Kakamyrat"
	first_map[2] = "Selbi"
	for i:=0; i<len(first_map); i++{
		fmt.Println("The key: ",i,"is ",first_map[i])
	}
}