package repository

import (
	"CleanArchitecture_SampleApp/domain"
)

//依存関係の逆転の法則
type UserRepository interface {
	Insert(user domain.User) error
	SelectByAuthToken(authToken string) (*domain.User, error)
	SelectByPrimaryKey(userID string) (*domain.User, error)
	UpdateByPrimaryKey(userID string, name string) error
}
