package model

type Topic struct {
	Tid       int `gorm:"column:tid; PRIMARY_KEY"`
	TopicName string
	ParentId  int
}
