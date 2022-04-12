package model

import "time"

type QuestionLike struct {
	Uid  int       `gorm:"column:uid;"`
	Qid  int       `gorm:"column:qid;"`
	Time time.Time `gorm:"autoCreateTime"`
}
