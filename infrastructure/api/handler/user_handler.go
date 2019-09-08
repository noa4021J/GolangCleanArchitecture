package handler

import (
	"CleanArchitecture_SampleApp/infrastructure/api/dcontext"
	"CleanArchitecture_SampleApp/infrastructure/datastore"
	"CleanArchitecture_SampleApp/infrastructure/server/response"
	"CleanArchitecture_SampleApp/interface/controllers"
	"CleanArchitecture_SampleApp/interface/database"
	"CleanArchitecture_SampleApp/usecase/service"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type userHandler struct {
	useController controllers.UserController
}

type UserHandler interface {
	GetUser(writer http.ResponseWriter, request *http.Request)
	UpdateUser(writer http.ResponseWriter, request *http.Request)
}

func NewUserHandler(db *datastore.ConnectedSql) UserHandler {
	//DBの注入
	return &userHandler{
		useController: controllers.NewUserController(
			service.NewUserService(
				database.NewUserRepository(db),
			),
		),
	}
}

func (uh *userHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	// Contextから認証済みのユーザIDを取得
	ctx := request.Context()
	userID := dcontext.GetUserIDFromContext(ctx)
	if len(userID) == 0 {
		log.Println(errors.New("userID is empty"))
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	userGetResponse, err := uh.useController.GetUser(&userID)
	if err != nil {
		log.Println(errors.New("userID is not Found"))
		return
	}
	response.Success(writer, userGetResponse)
}

func (uh *userHandler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	// リクエストBodyから更新後情報を取得
	var userUpdateRequest controllers.UserUpdateRequest
	err := json.NewDecoder(request.Body).Decode(&userUpdateRequest)
	if err != nil {
		log.Printf("%+v\n", err)
		response.BadRequest(writer, "Invalid Request")
		return
	}

	// Contextから認証済みのユーザIDを取得
	ctx := request.Context()
	userID := dcontext.GetUserIDFromContext(ctx)
	if len(userID) == 0 {
		log.Println(errors.New("userID is empty"))
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	err = uh.useController.UpdateUser(&userID, &userUpdateRequest)
	if err != nil {
		log.Println(errors.New("userID is not Found"))
		return
	}
	response.Success(writer, http.StatusOK)
}
