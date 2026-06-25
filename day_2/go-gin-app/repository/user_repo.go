package repository

import (
	"example.com/gin-app/config"
    "example.com/gin-app/models"
)

func CreateUser(user models.User) error{
	query := "Insert Into users(name, email) values($1, $2)"
	_, err := config.DB.Exec(query, user.Name, user.Email)
	return err
}

func GetUsers() ([]models.User, error){
	rows, err := config.DB.Query("Select id, name, email FROM users")
	if err != nil{
		return nil, err
	}
	var users[] models.User
	for rows.Next(){
		var u  models.User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}
	 return users, nil
}