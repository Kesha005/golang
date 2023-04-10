package main


import ("fmt")

type rectangle struct{
	a float64 
	b float64
}


func (rect *rectangle)double(double_value int){
	rect.a = rect.a *  float64(double_value)
	rect.b = rect.b * float64(double_value)
}

func (rect rectangle) area()( float64){
	return rect.a * rect.b
}


func main(){
	rect := rectangle {2,3}
	rect.double(5)
	fmt.Println( rect.area())
	
}