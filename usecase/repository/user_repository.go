package repository

import (
	"CleanArchitecture_SampleApp/domain"
)

//依存関係の逆転の法則
type UserRepository interface {
	Store(user domain.User) error
	FindByAuthToken(authToken string) (*domain.User, error)
	FindByUserID(userID string) (*domain.User, error)
	UpdateByUserID(userID string, name string) error
}
