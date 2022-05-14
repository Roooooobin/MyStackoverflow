package cronjob

import (
	"MyStackoverflow/dao"
	"MyStackoverflow/dao/answersdao"
	"MyStackoverflow/dao/questionsdao"
	"MyStackoverflow/dao/usersdao"
	"MyStackoverflow/model"
	"time"
)

func getStatusFromCriteria(points int) string {
	if points <= 200 {
		return "basic"
	} else if points <= 1000 {
		return "intermediate"
	} else if points <= 3000 {
		return "advanced"
	} else if points <= 10000 {
		return "expert"
	} else {
		return "master"
	}
}

func getPointsForUser(uid int) int {

	nowTime := time.Now()
	answerLikes := 0
	bestAnswers := 0
	answers, err := answersdao.List("uid = ?", uid)
	if err != nil {
		return -1
	}
	for _, answer := range answers {
		answerLikes += answer.Likes
		// within a month, double likes score
		if answer.Time.After(nowTime.AddDate(0, -1, 0)) {
			answerLikes += answer.Likes
		}
		if answer.IsBest == byte(model.IsBest) {
			bestAnswers++
		}
	}
	questions, err := questionsdao.List("uid = ?", uid)
	if err != nil {
		return -1
	}
	questionLikes := 0
	for _, question := range questions {
		questionLikes += question.Likes
		// within a month, double likes score
		if question.Time.After(nowTime.AddDate(0, -1, 0)) {
			questionLikes += question.Likes
		}
	}
	points := answerLikes + 100*bestAnswers + questionLikes
	return points
}

func refresh() {
	users := make([]model.User, 0)
	err := dao.MyDB.Table(usersdao.TableUsers).Find(&users).Error
	if err != nil {
		return
	}
	for _, user := range users {
		newStatus := getStatusFromCriteria(getPointsForUser(user.Uid))
		if newStatus == user.Status {
			continue
		}
		updateMap := map[string]interface{}{
			"status": newStatus,
		}
		_ = usersdao.Update(updateMap, "uid = ?", user.Uid)
	}
}
