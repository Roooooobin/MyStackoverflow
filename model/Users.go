package model

type User struct {
	Uid      int `gorm:"column:uid;"`
	Username string
	Status   string `gorm:"default:Basic;"`
	Email    string
	Password string
	City     string
	State    string
	Country  string
	Profile  string
}
