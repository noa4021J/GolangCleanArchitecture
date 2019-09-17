package controllers

import (
	"CleanArchitecture_SampleApp/interface/database"
	"CleanArchitecture_SampleApp/interface/dcontext"
	"CleanArchitecture_SampleApp/interface/network"
	"CleanArchitecture_SampleApp/usecase/service"
	"encoding/json"
	"errors"
	"log"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	GetUser(network.ApiResponser)
	UpdateUser(network.ApiResponser)
}

func NewUserController(db database.ConnectedSql) UserController {
	return &userController{
		userService: service.NewUserService(
			database.NewUserRepository(db),
		),
	}
}

func (uc *userController) GetUser(ar network.ApiResponser) {
	// Contextから認証済みのユーザIDを取得
	ctx := ar.GetRequestContext()
	userID := dcontext.GetUserIDFromContext(ctx)
	if len(userID) == 0 {
		log.Println(errors.New("userID is empty"))
		ar.InternalServerError("Internal Server Error")
		return
	}

	user, err := uc.userService.GetUser(&userID)
	if err != nil {
		return
	}

	userGetResponse := UserGetResponse{
		Name: user.Name,
	}

	ar.Success(userGetResponse)
}

func (uc *userController) UpdateUser(ar network.ApiResponser) {

	var userUpdateRequest UserUpdateRequest
	err := json.NewDecoder(ar.GetRequest().GetBody()).Decode(&userUpdateRequest)
	if err != nil {
		log.Printf("%+v\n", err)
		ar.BadRequest("Invalid Request")
		return
	}

	// Contextから認証済みのユーザIDを取得
	ctx := ar.GetRequestContext()
	userID := dcontext.GetUserIDFromContext(ctx)
	if len(userID) == 0 {
		log.Println(errors.New("userID is empty"))
		ar.InternalServerError("Internal Server Error")
		return
	}

	err = uc.userService.UpdateUser(&userID, &userUpdateRequest.Name)
	if err != nil {
		return
	}

	ar.Success(200)
}

type UserGetResponse struct {
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}
