package nosql

import (
	"fmt"
	"oop/model"
)

type noSqlUserRepository struct {
}

func NewNoSqlUserRepository() *noSqlUserRepository {
	return &noSqlUserRepository{}
}

func (n *noSqlUserRepository) InsertUser(user *model.User) error {
	fmt.Println("Inserting user on a NOSQL database")

	return nil
}
