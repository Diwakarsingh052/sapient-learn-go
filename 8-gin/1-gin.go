package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//r := gin.Default()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())
	
	r.GET("/json", sendJson)
	// Route Parameters
	// name is going to be a parameter
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name") // fetching the param
		c.String(200, "Hello, %s! (Gin)", name)
	})

	// Query Parameters
	r.GET("/welcome", func(c *gin.Context) {
		// this is setting the default value if not set
		firstName := c.DefaultQuery("firstName", "Guest")
		lastName := c.Query("lastName") // fetching the query
		c.String(200, "Hello, %s %s! (Gin)", firstName, lastName)
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			//panic("something went wrong")
			c.String(200, "Users v1 (Gin)")
		})
		v1.POST("/posts", func(c *gin.Context) {
			c.String(200, "Posts v1 (Gin)")
		})
	}

	r.Run(":8080")
}

func sendJson(c *gin.Context) {

	// setting the status code
	// setting the content type
	// converting the json
	//sending the response
	s := struct{ msg string }{msg: "hello world"}
	c.JSON(200, s)

}
