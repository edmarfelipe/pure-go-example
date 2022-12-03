package main

import (
	"log"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetUserInfo(username string) (result *UserInfo, err error) {
	defer func(start time.Time) {
		log.Printf("result=%v err=%v took=%v \n", result, err, time.Since(start))
	}(time.Now())

	return s.next.GetUserInfo(username)
}
