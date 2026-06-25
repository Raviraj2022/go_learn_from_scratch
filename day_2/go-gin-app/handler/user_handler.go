package handler


import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"example.com/gin-app/models"
	"example.com/gin-app/service"
)

func CreateUser(c *gin.Context) {
	var user models.User
 
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	err:= service.CreateUser(user)
	if err != nil {
		fmt.Println("DB Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB Error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context){
	users, err := service.GetUsers()

	if err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed"})
		return
	}
	c.JSON(http.StatusOK, users)
}