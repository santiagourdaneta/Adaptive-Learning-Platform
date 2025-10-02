package main

import (
	"github.com/gin-gonic/gin"
)

// setupRoutes configures all the API routes.
func setupRoutes(r *gin.Engine) { 
	r.POST("/register", func(c *gin.Context) { registerUserHandler(c) })
	r.GET("/path/:username", func(c *gin.Context) { getLearningPath(c) })
	r.GET("/content/:topicID", func(c *gin.Context) { getContentByTopic(c) })
	r.POST("/answer", func(c *gin.Context) { submitAnswer(c) })
	r.GET("/search", func(c *gin.Context) { searchContent(c) })
}