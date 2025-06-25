package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// https://github.com/gin-contrib/cors?tab=readme-ov-file#examples
type User struct {
	ID        uint      `json:"id" binding:"required"`
	Name      string    `json:"name" binding:"required,min=3,max=50"`
	Email     string    `json:"email" binding:"required,email"`
	Age       int       `json:"age" binding:"required,gte=18,lte=120"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	r := gin.Default()

	// this would apply middleware to all the routes
	r.Use(gin.Recovery())
	r.POST("/users", createUser)
	r.GET("/gin", func(c *gin.Context) {
		c.String(http.StatusOK, "Users (Gin)")
	})

	v1 := r.Group("/v1")
	{
		// this applies to v1 routes registered below
		v1.Use(SpecialLogger())
		v1.GET("/users", func(c *gin.Context) {
			//panic("something went wrong")
			c.String(200, "Users v1 (Gin)")
		})

		// auth would be applied to below routes only not the routes above
		v1.Use(Auth())
		v1.GET("/posts", func(c *gin.Context) {
			c.String(200, "Posts v1 (Gin)")
		})
	}

	v2 := r.Group("/v2")
	{
		v2.Use(Auth())
		v2.GET("/users", func(c *gin.Context) {
			//panic("something went wrong")
			c.String(200, "Users v1 (Gin)")
		})
		v2.POST("/posts", func(c *gin.Context) {
			c.String(200, "Posts v1 (Gin)")
		})
	}

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

// SpecialLogger middleware that can be used with r.Use()
func SpecialLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Time before request
		t := time.Now()

		// Process request
		c.Next()

		// After request
		latency := time.Since(t)
		log.Printf("Path: %s | Latency: %v | Status: %d ",
			c.Request.URL.Path,
			latency,
			c.Writer.Status(),
		)
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("auth mid layer started")
		c.Next()
	}
}
