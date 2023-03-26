package main

import (
	"fmt"
)

func main(){
	x := [4]int{1,2,3,4}
	t := x[1:3]
	fmt.Println(t)
}