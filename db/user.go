package db

import (
	"encoding/json"

	"github.com/ake-persson/mapslice-json"
)

type RemoteUser struct {
	Acct string `gorm:"acct"`
	Id   string `gorm:"id"`
	Raw  string `gorm:"raw"`
}

func CreateUserAS(
	username string, password string, privateKey string,
) (
	string, error,
) {
	ms := mapslice.MapSlice{
		mapslice.MapItem{Key: "abc", Value: 123},
		mapslice.MapItem{Key: "def", Value: 456},
		mapslice.MapItem{Key: "ghi", Value: 789},
	}

	b, err := json.Marshal(ms)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

type User struct {
	Username   string `gorm:"username"`
	Password   string `gorm:"password"`
	Email      string `gorm:"email"`
	PrivateKey string `gorm:"private_key"`
}
