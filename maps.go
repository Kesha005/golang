package main 

import (
	"fmt"
)


func  main(){
	fmt.Println("It's maps")

	m_a_p := map[int]string{
        90: "Dog",
        91: "Cat",
        92: "Cow",
        93: "Bird",
        94: "Rabbit",
    }
 
	fmt.Println("And this is a map" + m_a_p[90])
}