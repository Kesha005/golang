package main

import (
	"fmt"
	"os"
)


func CreateFile(filename string){
	file, err:= os.Create(filename);
	if err!=nil{
		fmt.Println(err)
		return 
	}

	defer file.Close()

	_,Writeerr:= file.WriteString("Hey ist new file Kerim")
	if Writeerr!=nil{
		fmt.Println(Writeerr.Error())
		return 
	}
}


func  WriteBytes(filename string, data []byte){
	file,err := os.Create(filename)
	if err!=nil {
		fmt.Println(err.Error())
		return 
	}
	defer file.Close()
	size, Wrerror := file.Write(data)
	if Wrerror!=nil{
		fmt.Println(Wrerror.Error())
		return 
	}
	fmt.Printf("Wrote %d of data to file", size )
}


func WritebyLine(filename string, lines []string){
	file, err:= os.Create(filename)
	
	if err!=nil {
		fmt.Println(err.Error())
		return 
	}
	defer file.Close()
	for _, val := range lines {
		_,err := fmt.Fprintf(file,"%s\n",val)
		if err!=nil{
			fmt.Println(err.Error())
			return 
		}

	}
}


func DelFile(filename string){
	err := os.Remove(filename)
	if err!=nil{
		fmt.Println(err.Error())
		return 
	}
}

func main(){
	// CreateFile("test.txt")
	// WriteBytes("second.txt",[]byte("Hello I'm a byte size"))
	lines := []string{
		"This is firts",
		"This is second",
		"This is third",
	};

	WritebyLine("thirt.txt",lines)
	DelFile("thiwrt.txt")
}