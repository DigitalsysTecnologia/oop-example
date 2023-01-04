package mocks

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"oop/model"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) InsertUser(user *model.User) error {
	args := m.Called(user)
	fmt.Println("Calling the mock object")
	return args.Error(0)
}
