package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	ch := make(chan string)

	from := "this must be email"
	pass := "this must be app password"

	host := "smtp.gmail.com"
	port := "587"
	to := []string{"gordonfireman1@gmail.com"}
	body := "Hello from golang"
	msg := []byte(body)
	auth := smtp.PlainAuth("", from, pass, host)
	for i := 0; i < 6; i++ {
		go func() {
			err := smtp.SendMail(host+":"+port, auth, from, to, msg)
			if err != nil {
				fmt.Print(err)
				ch <- err.Error()
			}
			ch <- "Mail sended successfully"
		}()
	}

	fmt.Println("mail sending in process")
	fmt.Println(<-ch)

}
