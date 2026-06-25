package routes

import (
	"github.com/gin-gonic/gin"
	"example.com/gin-app/handler"
)


func SetupRoutes(r *gin.Engine){
      r.POST("/users", handler.CreateUser)
	  r.GET("/users", handler.GetUsers)
}