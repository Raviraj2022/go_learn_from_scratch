package main

import "fmt"

// define struct
type User struct {
    Name string
    Age  int
}

func main() {
    // create struct
    user := User{
        Name: "Ravi",
        Age:  25,
    }

    fmt.Println(user)
}