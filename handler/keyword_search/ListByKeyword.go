package keyword_search

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/dao/usersdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func getUidSetFromUsernames(onlyUser string) map[int]struct{} {

	uidSet := make(map[int]struct{})
	if function.CheckNotEmpty(onlyUser) {
		usernames := strings.Split(onlyUser, ",")
		sqlStatement := "select uid from Users where username in (?)"
		rows, err := dao.MyDB.Table(usersdao.TableUsers).Raw(sqlStatement, usernames).Rows()
		if err != nil {
			return map[int]struct{}{}
		}
		for rows.Next() {
			var uid int
			err = rows.Scan(&uid)
			if err != nil {
				fmt.Println(err)
				return map[int]struct{}{}
			}
			uidSet[uid] = struct{}{}
		}
	}
	return uidSet
}

// extra options to filter questions, return true if passed all filters
func questionFilters(question *model.Question, isResolved, likesStr string, userSet map[int]struct{}) bool {

	if function.CheckNotEmpty(isResolved) && question.IsResolved != 1 {
		return false
	}
	// the question is not posted by any user you look for, return false
	if len(userSet) != 0 {
		_, ok := userSet[question.Uid]
		if !ok {
			return false
		}
	}
	if function.CheckNotEmpty(likesStr) {
		likes, err := strconv.Atoi(likesStr)
		if err == nil {
			if question.Likes < likes {
				return false
			}
		}
	}
	return true
}

// extra options to filter answers, return true if passed all filters
func answerFilters(answer *model.Answer, isBest, likesStr string, userSet map[int]struct{}) bool {

	if function.CheckNotEmpty(isBest) && answer.IsBest != 1 {
		return false
	}
	// the answer is not posted by any user you look for, return false
	if len(userSet) != 0 {
		_, ok := userSet[answer.Uid]
		if !ok {
			return false
		}
	}
	if function.CheckNotEmpty(likesStr) {
		likes, err := strconv.Atoi(likesStr)
		if err == nil {
			if answer.Likes < likes {
				return false
			}
		}
	}
	return true
}

func ListByKeyword(c *gin.Context) {
	/*
		list questions / answers / both by keyword
	*/
	errMsg := ""
	defer func() {
		if errMsg != "" {
			c.JSON(common.ErrorStatusCode, errMsg)
		}
	}()
	keyword := c.Query("keyword")
	if !function.CheckNotEmpty(keyword) {
		errMsg = "Input keyword can not be empty, please check and retry."
		return
	}
	sortByTime := c.Query("sortByTime")
	sortByLikes := c.Query("sortByLikes")
	// filter questions that are resolved
	isResolved := c.Query("isResolved")
	questionOnlyUsers := c.Query("questionOnlyUsers")
	questionUidSet := getUidSetFromUsernames(questionOnlyUsers)
	questionLikes := c.Query("questionLikes")
	// filter answers that are best
	isBest := c.Query("isBest")
	// filter answers only posted by this(these) user(s) with usernames, separated by ','
	answerOnlyUsers := c.Query("answerOnlyUsers")
	answerUidSet := getUidSetFromUsernames(answerOnlyUsers)
	answerLikes := c.Query("answerLikes")
	data := make(map[string]interface{})
	_, ok := c.GetQuery("onlyAnswers")
	if !ok {
		questionScoreMap := function.CalculateRelevanceScoreForQuestion(keyword)
		sql := dao.MyDB.Table(questionsdao.TableQuestions)
		questionsAll := make([]*model.Question, 0)
		err := sql.Find(&questionsAll).Error
		if err != nil {
			errMsg = err.Error()
			return
		}
		questions := make([]*model.Question, 0)
		for _, question := range questionsAll {
			if questionScoreMap[question.Qid] > 0 {
				if questionFilters(question, isResolved, questionLikes, questionUidSet) {
					questions = append(questions, question)
				}
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
			errMsg = err.Error()
			return
		}
		answers := make([]*model.Answer, 0)
		for _, answer := range answersAll {
			if answerScoreMap[answer.Aid] > 0 {
				if answerFilters(answer, isBest, answerLikes, answerUidSet) {
					answers = append(answers, answer)
				}
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
	if errMsg == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
