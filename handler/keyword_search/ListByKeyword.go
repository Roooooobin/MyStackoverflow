package keyword_search

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strings"
)

func ListByKeyword(c *gin.Context) {
	/*
		list questions / answers / both by keyword
	*/
	keyword := c.Query("keyword")
	if strings.Trim(keyword, " ") == "" {
		return
	}
	sortByTime := c.Query("sortByTime")
	sortByLikes := c.Query("sortByLikes")
	data := make(map[string]interface{})
	_, ok := c.GetQuery("onlyAnswers")
	if !ok {
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
		if sortByTime == "desc" {
			sort.Slice(questions, func(i, j int) bool {
				return questions[i].Time.After(questions[j].Time)
			})
		} else if sortByTime == "asc" {
			sort.Slice(questions, func(i, j int) bool {
				return questions[i].Time.Before(questions[j].Time)
			})
		}
		if sortByLikes == "desc" {
			sort.Slice(questions, func(i, j int) bool {
				return questions[i].Likes > questions[j].Likes
			})
		} else if sortByLikes == "asc" {
			sort.Slice(questions, func(i, j int) bool {
				return questions[i].Likes < questions[j].Likes
			})
		}
		data["questions"] = questions
	}
	_, ok = c.GetQuery("onlyQuestions")
	if !ok {
		answerScoreMap := function.CalculateRelevanceScoreForAnswer(keyword)
		fmt.Println(answerScoreMap)
		sql := dao.MyDB.Table(answersdao.TableAnswers)
		answersAll := make([]*model.Answer, 0)
		err := sql.Find(&answersAll).Error
		if err != nil {
			return
		}
		answers := make([]*model.Answer, 0)
		for _, answer := range answersAll {
			if answerScoreMap[answer.Aid] > 0 {
				answers = append(answers, answer)
			}
		}
		sort.Slice(answers, func(i, j int) bool {
			return answerScoreMap[answers[i].Aid] > answerScoreMap[answers[j].Aid]
		})
		if sortByTime == "desc" {
			sort.Slice(answers, func(i, j int) bool {
				return answers[i].Time.After(answers[j].Time)
			})
		} else if sortByTime == "asc" {
			sort.Slice(answers, func(i, j int) bool {
				return answers[i].Time.Before(answers[j].Time)
			})
		}
		if sortByLikes == "desc" {
			sort.Slice(answers, func(i, j int) bool {
				return answers[i].Likes > answers[j].Likes
			})
		} else if sortByLikes == "asc" {
			sort.Slice(answers, func(i, j int) bool {
				return answers[i].Likes < answers[j].Likes
			})
		}
		data["answers"] = answers
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
