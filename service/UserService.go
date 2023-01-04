package service

import (
	"errors"
	"oop/model"
)

type userRepository interface {
	InsertUser(user *model.User) error
}

type userService struct {
	userRepository userRepository
}

func NewUserService(userRepository userRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) InsertUser(user *model.User) error {
	err := u.userRepository.InsertUser(user)
	if err != nil {
		return errors.New("Error inserting user on service layer. " + err.Error())
	}

	return nil
}
