package main
import (
   "fmt"
   "os"    //import fmt and os package
)

func CreateDirectory(name string){
   if err := os.Mkdir(name,0750); err!=nil{//creates directory with name and permission code
      fmt.Println("There is Error "+err.Error())
   }

}


func main() {
   directory, err := os.Getwd()    //get the current directory using the built-in function
   if err != nil {
      fmt.Println(err) //print the error if obtained
   }
   fmt.Println("Current working directory:", directory) //print the required directory
   CreateDirectory("Kerim")
}