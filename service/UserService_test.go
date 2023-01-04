package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"oop/mocks"
	"oop/model"
	"testing"
)

func TestInsertUser(t *testing.T) {
	testCases := []struct {
		name             string
		userToBeInserted *model.User
		mockRepoError    error
		expectedError    error
	}{
		{
			name: "OK",
			userToBeInserted: &model.User{
				Name:     "Edu",
				Password: "123",
			},
			mockRepoError: nil,
			expectedError: nil,
		},
		{
			name: "Error inserting user",
			userToBeInserted: &model.User{
				Name:     "Wrong User",
				Password: "Wrong Password",
			},
			mockRepoError: errors.New("error inserting user"),
			expectedError: errors.New("Error inserting user on service layer. error inserting user"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			userRepository := &mocks.MockUserRepository{}
			userRepository.On("InsertUser", mock.Anything).Return(testCase.mockRepoError)

			userService := NewUserService(userRepository)
			err := userService.InsertUser(testCase.userToBeInserted)

			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
