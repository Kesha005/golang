package main

import "fmt"

func main() {
    ch := make(chan string, 1)

    // Start three goroutines that read from the channel
    for i := 0; i < 3; i++ {
        go func() {
            msg := <-ch
            fmt.Println("Received message:", msg)
        }()
    }

    // Broadcast a message to all the goroutines
    ch <- "Hello, World!"
}