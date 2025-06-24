package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        uint      `json:"id" binding:"required"`
	Name      string    `json:"name" binding:"required,min=3,max=50"`
	Email     string    `json:"email" binding:"required,email"`
	Age       int       `json:"age" binding:"required,gte=18,lte=120"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	r := gin.Default()
	r.POST("/users", createUser)
	r.Run(":8080")
}

func createUser(c *gin.Context) {
	var user User
	// validate the request
	// read the body
	// unmarshal json
	fmt.Println(c.Request.URL.Path)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.CreatedAt = time.Now()
	c.JSON(http.StatusOK, gin.H{"User created successfully": user})

}
