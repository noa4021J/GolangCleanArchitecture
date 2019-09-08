package handler

import (
	"CleanArchitecture_SampleApp/infrastructure/datastore"
	"CleanArchitecture_SampleApp/infrastructure/server/response"
	"CleanArchitecture_SampleApp/interface/controllers"
	"CleanArchitecture_SampleApp/interface/database"
	"CleanArchitecture_SampleApp/usecase/service"
	"encoding/json"
	"log"
	"net/http"
)

type authHandler struct {
	authController controllers.AuthController
}

type AuthHandler interface {
	CreateUser(writer http.ResponseWriter, request *http.Request)
}

func NewAuthHandler(db *datastore.ConnectedSql) AuthHandler {
	return &authHandler{
		authController: controllers.NewAuthController(
			service.NewAuthService(
				database.NewUserRepository(db),
			),
		),
	}
}

func (ah *authHandler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	// RequestBodyのパース
	var authCreateRequest controllers.AuthCreateRequest
	err := json.NewDecoder(request.Body).Decode(&authCreateRequest)
	if err != nil {
		log.Printf("%+v\n", err)
		response.BadRequest(writer, "Invalid Request")
		return
	}

	authCreateResponse, err := ah.authController.CreateUser(&authCreateRequest)
	if err != nil {
		response.InternalServerError(writer, "Internal Server Error")
		return
	}

	response.Success(writer, authCreateResponse)
}
