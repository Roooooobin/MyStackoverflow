package keyword_search

import (
	"MyStackoverflow/common"
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/answertopicsdao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/dao/questiontopicsdao"
	"MyStackoverflow/dao/usersdao"
	"MyStackoverflow/function"
	"MyStackoverflow/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func getUidsFromUsernames(onlyUser string) []int {

	uids := make([]int, 0)
	if function.CheckNotEmpty(onlyUser) {
		usernames := strings.Split(onlyUser, ",")
		sqlStatement := "select uid from Users where username in (?)"
		rows, err := dao.MyDB.Table(usersdao.TableUsers).Raw(sqlStatement, usernames).Rows()
		if err != nil {
			return nil
		}
		for rows.Next() {
			var uid int
			err = rows.Scan(&uid)
			if err != nil {
				return nil
			}
			uids = append(uids, uid)
		}
	}
	return uids
}

// extra options to filter questions, return true if passed all filters
func questionFilters(question *model.Question, likesStr string) bool {

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
	// usernames, separated by ','
	questionOnlyUsers := c.Query("questionOnlyUsers")
	questionUids := getUidsFromUsernames(questionOnlyUsers)
	questionLikes := c.Query("questionLikes")
	// filter answers that are best
	isBest := c.Query("isBest")
	// filter answers only posted by this(these) user(s) with usernames, separated by ','
	answerOnlyUsers := c.Query("answerOnlyUsers")
	answerUids := getUidsFromUsernames(answerOnlyUsers)
	answerLikes := c.Query("answerLikes")
	// topic ids, separated by ','
	topicStr := c.Query("topic")
	data := make(map[string]interface{})
	_, ok := c.GetQuery("onlyAnswers")
	if !ok {
		sql := dao.MyDB.Table(questionsdao.TableQuestions)
		// filters: is_resolved, uid, likes, topic
		if function.CheckNotEmpty(isResolved) {
			sql = sql.Where("is_resolved = ?", model.IsResolved)
		}
		if questionUids != nil {
			sql = sql.Where("uid in (?)", questionUids)
		}
		if function.CheckNotEmpty(questionLikes) {
			likes, err := strconv.Atoi(questionLikes)
			if err != nil {
				errMsg = "Question likes error: " + err.Error()
			}
			sql = sql.Where("likes >= ?", likes)
		}
		questionsCandidates := make([]*model.Question, 0)
		err := sql.Find(&questionsCandidates).Error
		if err != nil {
			errMsg = err.Error()
			return
		}
		questionScoreMap := function.CalculateRelevanceScoreForQuestion(keyword)
		questions := make([]*model.Question, 0)
		for _, question := range questionsCandidates {
			if questionScoreMap[question.Qid] > 0 {
				questions = append(questions, question)
			}
		}
		// get all questions belong to the topic
		questionInTopicSet := make(map[int]struct{})
		// filter by topic
		if function.CheckNotEmpty(topicStr) {
			topicIDs := strings.Split(topicStr, ",")
			candidateQids := make([]int, 0)
			questionsCandidates = questions
			for _, candidate := range questionsCandidates {
				candidateQids = append(candidateQids, candidate.Qid)
			}
			questionTopics, err := questiontopicsdao.List("qid in (?) and tid in (?)", candidateQids, topicIDs)
			if err != nil {
				errMsg = err.Error()
				return
			}
			for _, questionTopic := range questionTopics {
				questionInTopicSet[questionTopic.Qid] = struct{}{}
			}
			questions = make([]*model.Question, 0)
			for _, question := range questionsCandidates {
				if _, ok = questionInTopicSet[question.Qid]; ok {
					questions = append(questions, question)
				}
			}
		}
		// sort by relevance score
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
		sql := dao.MyDB.Table(answersdao.TableAnswers)
		// filters: is_best, uid, likes, topic
		if function.CheckNotEmpty(isBest) {
			sql = sql.Where("is_best = ?", model.IsBest)
		}
		if answerUids != nil {
			sql = sql.Where("uid in (?)", answerUids)
		}
		if function.CheckNotEmpty(answerLikes) {
			likes, err := strconv.Atoi(answerLikes)
			if err != nil {
				errMsg = "Answer likes error: " + err.Error()
			}
			sql = sql.Where("likes >= ?", likes)
		}
		answerCandidates := make([]*model.Answer, 0)
		err := sql.Find(&answerCandidates).Error
		if err != nil {
			errMsg = err.Error()
			return
		}
		answers := make([]*model.Answer, 0)
		// filter out answers whose relevance score = 0
		for _, answer := range answerCandidates {
			if answerScoreMap[answer.Aid] > 0 {
				answers = append(answers, answer)
			}
		}
		// get all questions belong to the topic
		answerInTopicSet := make(map[int]struct{})
		// filter by topic
		if function.CheckNotEmpty(topicStr) {
			topicIDs := strings.Split(topicStr, ",")
			candidateAids := make([]int, 0)
			answerCandidates = answers
			for _, candidate := range answerCandidates {
				candidateAids = append(candidateAids, candidate.Aid)
			}
			answerTopics, err := answertopicsdao.List("aid in (?) and tid in (?)", candidateAids, topicIDs)
			if err != nil {
				errMsg = err.Error()
				return
			}
			for _, answerTopic := range answerTopics {
				answerInTopicSet[answerTopic.Aid] = struct{}{}
			}
			answers = make([]*model.Answer, 0)
			for _, answer := range answerCandidates {
				if _, ok = answerInTopicSet[answer.Aid]; ok {
					answers = append(answers, answer)
				}
			}
		}
		// sort by relevance score
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
