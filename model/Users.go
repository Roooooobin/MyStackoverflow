package model

type User struct {
	Uid      int `gorm:"column:uid; primaryKey"`
	Username string
	Status   string `gorm:"default:basic;"`
	Email    string
	Password string
	City     string
	State    string
	Country  string
	Profile  string
}
