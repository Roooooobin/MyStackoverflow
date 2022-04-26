package model

import "time"

type Answer struct {
	Aid    int `gorm:"column:aid; primaryKey"`
	Qid    int `gorm:"column:qid"`
	Uid    int `gorm:"column:uid"`
	Body   string
	Time   time.Time `gorm:"autoCreateTime"`
	IsBest byte
	Likes  int
	Rating byte `gorm:"default:null"`
}
