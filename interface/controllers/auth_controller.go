package controllers

import (
	"CleanArchitecture_SampleApp/usecase/service"
)

type authController struct {
	authService service.AuthService
}

type AuthController interface {
	CreateUser(newUser *AuthCreateRequest) (*AuthCreateResponse, error)
}

func NewAuthController(as service.AuthService) AuthController {
	return &authController{authService: as}
}

//CreateUser
func (ac *authController) CreateUser(newUser *AuthCreateRequest) (*AuthCreateResponse, error) {
	authToken, err := ac.authService.CreateUser(&newUser.Name)
	if err != nil {
		return nil, err
	}
	return &AuthCreateResponse{Token: *authToken}, nil
}

type AuthCreateRequest struct {
	Name string `json:"name"`
}

type AuthCreateResponse struct {
	Token string `json:"token"`
}
