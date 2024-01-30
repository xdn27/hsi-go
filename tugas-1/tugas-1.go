package main

import "fmt"
import "errors"


func main() {
    fmt.Println("Hello, world! dengan Println")
    
    output := fmt.Sprintln("Hello, world! dengan Sprintln")
    fmt.Println(output)
    
    
    err := fmt.Errorf("Terjadi error: %s", "String error")
    fmt.Println(err)
    
    err2 := errors.New("Unexpected error")
    fmt.Println(err2)

}
