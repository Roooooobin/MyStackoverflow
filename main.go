package main

import (
	"MyStackoverflow/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	groupUser := r.Group("/user")
	{
		groupUser.POST("/add", func(c *gin.Context) {
			handler.AddUser(c)
		})
	}

	groupTopic := r.Group("/topic")
	{
		groupTopic.POST("/add", func(c *gin.Context) {
			handler.AddTopic(c)
		})
	}

	groupQuestion := r.Group("/question")
	{
		groupQuestion.POST("/add", func(c *gin.Context) {
			handler.AddQuestion(c)
		})
	}
	// listen and serve on 0.0.0.0:8080
	err := r.Run()
	if err != nil {
		return
	}
}
