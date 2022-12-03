package main

import (
	"net/http"
)

type APIServer struct {
	svc Service
}

func NewAPIServer(svc Service) *APIServer {
	return &APIServer{
		svc: svc,
	}
}

func (s *APIServer) Start(addr string) error {
	http.HandleFunc("/", makeHTTPHandler(s.handleGetUserInfo))
	return http.ListenAndServe(addr, nil)
}

func (s *APIServer) handleGetUserInfo(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return errMethodNotAllowed
	}

	info, err := s.svc.GetUserInfo("edmarfelipe")
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, info)
}
