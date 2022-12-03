package main

type MockUserService struct{}

func NewMockUserInfoService() Service {
	return &MockUserService{}
}

func (s *MockUserService) GetUserInfo(username string) (*UserInfo, error) {
	return &UserInfo{
		ID:        123,
		Name:      "My Name",
		Followers: 23,
		Following: 54,
	}, nil
}
