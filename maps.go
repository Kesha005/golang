package main 

import (
	"fmt"
)


func  main(){

    fmt.Println("It's a simple map")

    m_a_p := map[int]string{
        1:"Kerim",
        2:"Kakamyrat",
        3:"Selbi",
    }
    for d , i:=range m_a_p{
        fmt.Println("The index",d, "name",i)
    }
    
}

