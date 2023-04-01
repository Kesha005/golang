package main 

import(
	"fmt"
)


func main() {
    fmt.Println("Returned:", MyFunc())
}

func MyFunc() (ret string) {
    defer func() {
        if r := recover(); r != nil {
            ret = fmt.Sprintf("was panic, recovered value: %v", r)
        }
    }()
    panic("test")
    return "Normal Return Value"
}