package question

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

// ListQuestionByKeyword list questions by keyword
func ListQuestionByKeyword(c *gin.Context) {

	keyword := c.Query("keyword")
	questionScoreMap := function.CalculateRelevanceScoreForQuestion(keyword)
	sql := dao.MyDB.Table(questionsdao.TableQuestions)
	questionsAll := make([]*model.Question, 0)
	err := sql.Find(&questionsAll).Error
	if err != nil {
		return
	}
	questions := make([]*model.Question, 0)
	for _, question := range questionsAll {
		if questionScoreMap[question.Qid] > 0 {
			questions = append(questions, question)
		}
	}
	sort.Slice(questions, func(i, j int) bool {
		return questionScoreMap[questions[i].Qid] > questionScoreMap[questions[j].Qid]
	})
	c.JSON(http.StatusOK, gin.H{
		"data": questions,
	})
}
