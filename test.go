package main


import (
	"github.com/Kesha005/go_encryptor"
	"fmt"
	"github.com/joho/godotenv"

)


func main(){
	godotenv.Load(".env")
	StringToEncrypt := "Encrypting this string"
	fmt.Println(StringToEncrypt)
	encText, err := go_encryptor.Encrypt(StringToEncrypt)
	if err!=nil{
		panic(err)
	}
	fmt.Println(encText)
}