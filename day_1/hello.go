package main

import (
    "fmt"
    "example.com/greetings/models"
)

func main() {
    user := models.User{
        Name: "Ravi",
        Age:  25,
    }

    fmt.Println(user)
    // fmt.Println(user.Age)
}