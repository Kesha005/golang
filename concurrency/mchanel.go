package main


import "fmt"



func div(x,y int, c chan int)float64{
	result := x/y
	return float64(result)
}


func power(x,y int, c1 chan int)int{
	rslt := x ^ y
	return rslt
}


func main(){
	fmt.Println("Enter two numbers to operate: ")
	var x, y int
	var c, c1 chan int
	fmt.Scanln(&y)
	fmt.Scanln(&x)

	go div(x, y,c)
	go power(x,y, c1)

}