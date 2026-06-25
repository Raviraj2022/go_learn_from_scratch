package service


import (
	"example.com/gin-app/models"

	"example.com/gin-app/repository"
)

func CreateUser(user models.User)error {
	return repository.CreateUser(user)
}

func GetUsers()([]models.User, error){
	return repository.GetUsers()
}