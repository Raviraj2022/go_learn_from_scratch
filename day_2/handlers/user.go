package handlers

import (
	"encoding/json"
	 "net/http" 
	 "strconv"


"example.com/crud_app/models"
)

var  users []models.User
var idCounter = 1

func CreateUser( w http.ResponseWriter, r *http.Request){
	var user  models.User
	json.NewDecoder(r.Body).Decode(&user)

	user.ID = idCounter
    idCounter++

	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}


func GetAllUsers(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(users)
}

func GetSingleUsers(w http.ResponseWriter, r *http.Request){
	idParam := r.URL.Query().Get("id")
	id, _:= strconv.Atoi(idParam)

	for _, u := range users{
		if u.ID == id {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "User Not Found", http.StatusNotFound)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)

	var updated models.User
	json.NewDecoder(r.Body).Decode(&updated)

	for i, u := range users {
		if u.ID == id{
           users[i].Name = updated.Name
           users[i].Email = updated.Email

		   json.NewEncoder(w).Encode(users[i])
		   return
		}
	}
	http.Error(w, "user not found",  http.StatusNotFound )
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
       idParam := r.URL.Query().Get("id")
	   id, _ := strconv.Atoi(idParam)

	   for i, u := range users{
		if u.ID == id {
			users = append(users[:i], users[i + 1:]...)
			json.NewEncoder(w).Encode("Deleted")
			return
		}
	   }

	   http.Error(w, "User Not Found", http.StatusNotFound)
}