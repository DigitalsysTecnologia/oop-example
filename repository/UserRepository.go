package repository

import (
	"fmt"
	"oop/model"
)

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (u *userRepository) InsertUser(user *model.User) error {
	// Insert on database
	fmt.Println("Inserted user on database")

	return nil
}
