package main

type UserInfo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}
