package main 

import ( 
	"fmt"
 "net/http" 
 "example.com/crud_app/handlers"
)

func main(){
	http.HandleFunc("/users", handlers.GetAllUsers)
	http.HandleFunc("/user", handlers.GetSingleUsers)
	http.HandleFunc("/create", handlers.CreateUser)
	http.HandleFunc("/update", handlers.UpdateUser)
	http.HandleFunc("/delete", handlers.DeleteUser)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}