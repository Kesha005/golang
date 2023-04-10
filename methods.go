package main

import (
	"fmt"
)

type author struct {
	name       string
	surname    string
	birth_year int
}


func (book_author author) show(){
	fmt.Println("Author name is ", book_author.name);
	fmt.Println("Author surname is " , book_author.surname)
	fmt.Println("Author birth year is ", book_author.birth_year)
}


func main() {
	book_author := author {"Kerimberdi" , "Saparow", 2005}
	book_author.show()
}
