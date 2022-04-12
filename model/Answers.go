package model

import "time"

type Answer struct {
	Aid    int `gorm:"column:aid"`
	Qid    int `gorm:"column:qid"`
	Uid    int `gorm:"column:uid"`
	Body   string
	Time   time.Time `gorm:"autoCreateTime"`
	IsBest byte
	Likes  int
}
