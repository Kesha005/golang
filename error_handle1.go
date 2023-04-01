package main

import(
	"errors"
	"fmt"
)



func findUserOnMap(id int )(error , string){
	m_a_p := map[int]string {
		1 : "Kerim",
		2 : "Selbi",
		3 : "Selim",
	}
	val , ok :=m_a_p[id] 
	if (ok){
		return nil, val
	}
	errors.New("Bu Id boyunca tapylmady")
	return errors.New("Bu Id boyunca tapylmady") , "Yok"
	
}

func findOnArray(name string)(error, string){
	
	arr1:= [...]string{"GFG", "gfg", "geeks",
	"GeeksforGeeks", "GEEK"}

	for i := range arr1{
		if arr1[i] == name{
			return nil,arr1[i]
		}
		return errors.New("Tapylmady"),"null"
	}
	return errors.New("Tapylmady"),"null"
}

func main(){
	err , user := findUserOnMap(32)
		if(err != nil){
			fmt.Println(err)
		}
		fmt.Println(user)
	err1, name :=findOnArray("Kerim")
	if(err1 != nil){
		fmt.Println(err1)
	}

	fmt.Println(name)
		
}




