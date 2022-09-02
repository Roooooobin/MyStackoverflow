package model

type UserStats struct {
	Uid             int `gorm:"column:uid; primaryKey"`
	QuestionsLength int
	QuestionsWords  int
	AnswersLength   int
	AnswersWords    int
}
