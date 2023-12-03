package main


import(
	"fmt"
	"os"
)




func WriteToFile(filename string, data string){

	file, err := os.Open(filename)
	if err!=nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	_, wrError := file.WriteString(data)
	if wrError!=nil{
		fmt.Println(wrError.Error())
		return
	}
}


func CreateNewFile(filename string){
	_,err:= os.Create(filename)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
}



func main(){
	CreateNewFile("fourth.txt")
	WriteToFile("fourth.txt","Hi Kerim this is writer")
}