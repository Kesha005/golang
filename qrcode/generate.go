package main


import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"sync"

)


func main(){
	urls := []string{
		"Kerim",
		"Kakamyrat",
		"Selbi",
	}
	var wg sync.WaitGroup

	wg.Add(len(urls))
	for i , url := range urls{
		url := url
		i := i
		go func(){
			defer wg.Done()
			GenerateQr(i,url)
		}()
	}
	wg.Wait()
}


func GenerateQr(i int, url string){
	qr, _ := qrcode.New(url, qrcode.Medium)
	filename := fmt.Sprintf("%v.png",i)
	err := qr.WriteFile(256,filename)
	if err!=nil{
		fmt.Println("There is some error in generating qr")
		return 
	}
	
	fmt.Sprintf("Qr code generated and saved %v.png", i)

}