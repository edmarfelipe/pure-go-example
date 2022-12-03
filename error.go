package main

import "net/http"

var (
	errMethodNotAllowed = apiError{Err: "Not Allowed", status: http.StatusMethodNotAllowed}
)

type apiError struct {
	Err    string `json:"error"`
	status int
}

func (e apiError) Error() string {
	return e.Err
}
