package main

import (
	"fmt"
	"net/http"
)



func main(){

	mychannel := make(chan string)

    links := []string{
		"https://github.com/Kesha005",
		"https://github.com",
	}

	for _ , link := range links{
		go getdatafromLink(link ,mychannel)
	}

	for i:=0 ; i< len(links); i++{
		fmt.Printf("Link %v is up \n", <- mychannel)
	}

}


func getdatafromLink(link string , mychannel chan string){
	_,err:=http.Get(link)
	if err ==nil{
		mychannel <- link
	}
}