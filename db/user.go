package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/Hana-ame/moonchan-backend/utils"
	"github.com/Hana-ame/moonchan-backend/webfinger"
	"github.com/ake-persson/mapslice-json"
	"github.com/valyala/fastjson"
)

type RemoteUser struct {
	Acct string `gorm:"acct"`
	Id   string `gorm:"id"`
	As   []byte `gorm:"as"`
}

func GetResourceByAcct(acct string) (*webfinger.Resource, error) {
	username, host := utils.ParseAcct(acct)
	url := fmt.Sprintf("https://%s/.well-known/webfinger?resource=acct:%s@%s", host, username, host)

	resp, err := utils.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resource webfinger.Resource
	if err := json.NewDecoder(resp.Body).Decode(&resource); err != nil {
		return nil, err
	}

	return &resource, nil
}

func GetRemoteUserAsByAcct(acct string) ([]byte, error) {
	resource, err := GetResourceByAcct(acct)
	if err != nil {
		return nil, err
	}

	id, err := webfinger.GetIdFromResource(resource)
	if err != nil {
		return nil, err
	}

	resp, err := utils.Get(id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func GetRemoteUserObjByAs(acct string, as []byte) (*RemoteUser, error) {
	if !fastjson.Exists(as, "id") {
		return nil, errors.New("not found id in as")

	}
	id := fastjson.GetString(as, "id")
	return &RemoteUser{
		Acct: acct,
		Id:   id,
		As:   as,
	}, nil
}

type LocalUser struct {
	Username   string `gorm:"username"`
	Password   string `gorm:"password"`
	Email      string `gorm:"email"`
	PrivateKey string `gorm:"private_key"`
	As         []byte `gorm:"as"`
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
