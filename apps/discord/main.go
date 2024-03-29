package main

import (
	"net/http"
	"strconv"

	"github.com/calamity-m/fern/middleware"
	"github.com/gin-gonic/gin"
)

type TestController struct {
	Id int
}

func (c *TestController) GetTestInstance(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "Non integer ID"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"New user": TestController{Id: id}})
}

func Server() {
	// Create a new Gin router
	router := gin.Default()
	tc := &TestController{}

	router.Use(middleware.Timer())
	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Fern": "Says hello!"})
	})

	test := router.Group("/test")
	{
		test.GET("/test/:id", tc.GetTestInstance)

		// Route with URL parameters
		test.GET("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.String(200, "User ID: "+id)
		})

		// Route with query parameters
		test.GET("/search", func(c *gin.Context) {
			query := c.DefaultQuery("q", "default-value")
			c.String(200, "Search query: "+query)
		})
	}

	// Run the server on port 8000
	router.Run(":8000")

}
