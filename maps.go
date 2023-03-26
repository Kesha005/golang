package main

import(
	"fmt"
)

func main(){
	maps :=make(map[string]string)
	maps["hello"]="salam"
	maps["goodbye"]="sagbol"
	if name ,ok :=maps["hello"]; ok {
		fmt.Println(name,ok)
	}
	
}