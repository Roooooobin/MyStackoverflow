package main

import (
	"MyStackoverflow/cache"
	"MyStackoverflow/handler/answer"
	"MyStackoverflow/handler/question"
	"MyStackoverflow/handler/topic"
	"MyStackoverflow/handler/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	groupUser := r.Group("/user")
	{
		groupUser.POST("/add", func(c *gin.Context) {
			user.AddUser(c)
		})
		groupUser.GET("/list", func(c *gin.Context) {
			user.ListUser(c)
		})
		groupUser.POST("/edit", func(c *gin.Context) {
			user.EditUser(c)
		})
	}

	groupTopic := r.Group("/topic")
	{
		groupTopic.POST("/add", func(c *gin.Context) {
			topic.AddTopic(c)
		})
	}

	groupQuestion := r.Group("/question")
	{
		groupQuestion.POST("/add", func(c *gin.Context) {
			question.AddQuestion(c)
		})
		groupQuestion.POST("/like", func(c *gin.Context) {
			question.LikeQuestion(c)
		})
		groupQuestion.GET("/list", func(c *gin.Context) {
			question.ListQuestion(c)
		})
		groupQuestion.GET("/listbykeyword", func(c *gin.Context) {
			question.ListQuestionByKeyword(c)
		})
	}

	groupAnswer := r.Group("/answer")
	{
		groupAnswer.POST("/add", func(c *gin.Context) {
			answer.AddAnswer(c)
		})
		groupAnswer.POST("/like", func(c *gin.Context) {
			answer.LikeAnswer(c)
		})
		groupAnswer.GET("/list", func(c *gin.Context) {
			answer.ListAnswer(c)
		})
		groupAnswer.POST("/select", func(c *gin.Context) {
			answer.SelectBest(c)
		})
		groupAnswer.POST("/rate", func(c *gin.Context) {
			answer.RateAnswer(c)
		})
	}

	// pre-computed cache
	cache.Init()
	// listen and serve on 0.0.0.0:8080
	err := r.Run()
	if err != nil {
		return
	}
}
