package main
import "fmt"

func main() {

  // create channel
  number := make(chan int)
  message := make(chan string)

  // function call with goroutine
  go chh(number, message)

  // retrieve channel data
  fmt.Println("Channel Data:", <-number)
  fmt.Println("Channel Data:", <-message)

}

func chh(number chan int, message chan string) {

  // send data into channel
  number <- 15
  message <- "Learning Go channel"

}