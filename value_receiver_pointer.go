package main


import("fmt")



type value struct{
	width float64
	height float64
}


func (object *value)add(adding int){
	object.width = object.width * float64(adding)
	object.height = object.height * float64(adding)
}


func power(){
	
}



func main(){
	object := &value{3,4}
	fmt.Printf("%+v\n", "before adding ", object)
	object.add(3)
	fmt.Println("after ", object)


}