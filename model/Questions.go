package model

import "time"

type Question struct {
	Qid        int `gorm:"column:qid;"`
	Uid        int `gorm:"column:uid;"`
	Title      string
	Body       string
	Time       time.Time `gorm:"autoCreateTime"`
	IsResolved byte      `gorm:"default:0"`
	BestAid    int       `gorm:"default:null"`
	Likes      int       `gorm:"default:0"`
}
