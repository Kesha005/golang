package main

import (
	"fmt"
)

func main(){

	type Person struct{
		name ,surname ,adress string
		birthyear int
	}

	a := Person{"Kerimberdi ", "Saparow","Mekan",2005}

	fmt.Println(a.name,a.surname," was born in ",a.birthyear,".And he lives in ",a.adress,"street
	")

}