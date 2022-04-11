package usersdao

import (
	"MyStackoverflow/model"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	user := model.User{
		Username: "Brice",
		Email:    "brice@gmail.com",
		Password: "123456Brice",
		City:     "Queen",
		State:    "NY",
		Country:  "US",
		Profile:  "I am Brice from Queen, 23, graduate",
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
