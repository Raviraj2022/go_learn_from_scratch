// package main

// import (
//     "github.com/gin-gonic/gin"
// )

// func main() {
//     r := gin.Default()

//     r.GET("/", func(c *gin.Context) {
//         c.JSON(200, gin.H{
//             "message": "Hello from Gin 🚀",
//         })
//     })

//     r.Run(":8080")
// }

package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "example.com/gin-app/config"
    "example.com/gin-app/routes"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env")
    }

    config.ConnectDB()

    r := gin.Default()
    routes.SetupRoutes(r)

    r.Run(":5080")
}