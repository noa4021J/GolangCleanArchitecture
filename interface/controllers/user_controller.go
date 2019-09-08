package controllers

import (
	"CleanArchitecture_SampleApp/usecase/service"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	GetUser(userID *string) (*UserGetResponse, error)
	UpdateUser(userID *string, updateRequest *UserUpdateRequest) error
}

func NewUserController(us service.UserService) UserController {
	return &userController{userService: us}
}

func (uc *userController) GetUser(userID *string) (*UserGetResponse, error) {
	user, err := uc.userService.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return &UserGetResponse{Name: user.Name}, nil
}

func (uc *userController) UpdateUser(userID *string, updateRequest *UserUpdateRequest) error {
	err := uc.userService.UpdateUser(userID, &updateRequest.Name)
	if err != nil {
		return err
	}
	return nil
}

type UserGetResponse struct {
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}
