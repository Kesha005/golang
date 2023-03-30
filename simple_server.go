package main

import(
	"fmt"
	"log"
	"net/http"
)


func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r * http.Request){
		fmt.Fprintf(w ,"Hello world")
	})
	port := ":8000"

	fmt.Println("Server is runnin on port" + port)

	log.Fatal(http.ListenAndServe(port,nil))
}