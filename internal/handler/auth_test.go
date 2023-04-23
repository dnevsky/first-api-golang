package handler

import (
	"testing"

	mock_service "github.com/dnevsky/firstapi/internal/service/mocks"
	"github.com/dnevsky/firstapi/models"
	"github.com/golang/mock/gomock"
)

// сюда НЕТ СМЫСЛА ЛЕЗТЬ
func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user models.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           models.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			inputBody: `{"name":"Test","username":"test","password:"qwerty"}`,
			inputUser: models.User{
				Name:     "Test",
				Username: "test",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)
		})
	}
}
