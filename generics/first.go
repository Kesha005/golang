package main


import ("fmt")



func second_func[v string | int](msg v){
	fmt.Println(msg)
}


func main(){
	second_func[int](3)
	second_func[string]("Kerim")
}