package service

import (
	"CleanArchitecture_SampleApp/domain"
	"CleanArchitecture_SampleApp/usecase/repository"
)

type userService struct {
	UserRepository repository.UserRepository
}

type UserService interface {
	GetUser(userID *string) (*domain.User, error)
	UpdateUser(userID, newName *string) error
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{UserRepository: ur}
}

func (userService *userService) GetUser(userID *string) (*domain.User, error) {

	user, err := userService.UserRepository.FindByUserID(*userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService *userService) UpdateUser(userID, newName *string) error {

	err := userService.UserRepository.UpdateByUserID(*userID, *newName)
	if err != nil {
		return err
	}

	_, err = userService.UserRepository.FindByUserID(*userID)
	if err != nil {
		return err
	}

	return nil
}
