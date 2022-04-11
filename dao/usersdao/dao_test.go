package usersdao

import (
	"MyStackoverflow/model"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	user := model.User{
		Username: "Chris",
		Email:    "chris@gmail.com",
		Password: "123456Chris",
		City:     "Manhattan",
		State:    "NY",
		Country:  "US",
		Profile:  "I am Chris, nothing else to say",
	}
	if err := Insert(user); err != nil {
		fmt.Println(err)
		return
	}
}

func TestFind(t *testing.T) {
	user, _ := Find("uid = ?", 2)
	fmt.Println(user.Username, user.Profile)
}
