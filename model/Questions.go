package model

import "time"

type Question struct {
	Qid        int `gorm:"column:qid;"`
	Uid        int `gorm:"column:uid;"`
	Title      string
	Body       string
	Time       time.Time `gorm:"autoCreateTime"`
	IsResolved byte
	BestAid    int `gorm:"default:null"`
}
