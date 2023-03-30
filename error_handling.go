package main 


import(
	"fmt"
	"math"
	"errors"
)



func Sqrt(value float64)(float64 ,error){
	if(value <0 ){
		return 0 ,errors.New("Division error")
	}
	return math.Sqrt(value),nil
}

func main(){
	result , err := Sqrt(-1)
	if(err !=nil){
		fmt.Println(err)
	}
	fmt.Println(result)
}