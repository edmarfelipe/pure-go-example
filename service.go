package main

import (
	"encoding/json"
	"net/http"
)

type Service interface {
	GetUserInfo(usename string) (*UserInfo, error)
}

type UserService struct {
	url string
}

func NewUserInfoService(url string) Service {
	return &UserService{
		url: url,
	}
}

func (s *UserService) GetUserInfo(username string) (*UserInfo, error) {
	resp, err := http.Get(s.url + username)
	if err != nil {
		return nil, err
	}

	info := &UserInfo{}
	if err := json.NewDecoder(resp.Body).Decode(info); err != nil {
		return nil, err
	}

	return info, nil
}
