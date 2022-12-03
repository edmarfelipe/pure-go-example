package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server = APIServer{
		svc: NewMockUserInfoService(),
	}
)

func TestGetUserInfo(t *testing.T) {
	testServer := httptest.NewServer(makeHTTPHandler(server.handleGetUserInfo))
	resp, err := http.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", resp.StatusCode)
	}

	result := UserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Error(err)
	}

	if result.Name != "My Name" {
		t.Errorf("expected My Name but got %s", result.Name)
	}
}

func TestGetUserInfoRR(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}

	server.handleGetUserInfo(rr, req)

	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", rr.Result().StatusCode)
	}

	result := UserInfo{}
	err = json.NewDecoder(rr.Result().Body).Decode(&result)
	if err != nil {
		t.Error(err)
	}

	if result.Name != "My Name" {
		t.Errorf("expected My Name but got %s", result.Name)
	}
}
