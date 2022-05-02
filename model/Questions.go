package model

import "time"

var IsResolved = 1

type Question struct {
	Qid        int `gorm:"column:qid; primaryKey"`
	Uid        int `gorm:"column:uid;"`
	Title      string
	Body       string
	Time       time.Time `gorm:"autoCreateTime"`
	IsResolved byte      `gorm:"default:0"`
	Likes      int       `gorm:"default:0"`
}

type QuestionWithAnswerNum struct {
	Qid         int `gorm:"column:qid; primaryKey"`
	Uid         int `gorm:"column:uid;"`
	Title       string
	Body        string
	Time        time.Time `gorm:"autoCreateTime"`
	IsResolved  byte      `gorm:"default:0"`
	Likes       int       `gorm:"default:0"`
	NumOfAnswer int
}
