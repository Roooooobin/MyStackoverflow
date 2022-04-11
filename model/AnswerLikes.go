package model

import (
	"time"
)

type AnswerLike struct {
	Uid  int       `gorm:"column:uid;"`
	Aid  int       `gorm:"column:aid;"`
	Time time.Time `gorm:"autoCreateTime"`
}
