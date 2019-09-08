package handler

import (
	"CleanArchitecture_SampleApp/infrastructure/datastore"
	"net/http"
)

type interactor struct {
	db *datastore.ConnectedSql
}

type Interactor interface {
	NewAppHandler() AppHandler
}

func NewInteractor(db *datastore.ConnectedSql) Interactor {
	return &interactor{db: db}
}

func (i *interactor) NewAppHandler() AppHandler {
	return &appHandler{
		authHandler: NewAuthHandler(i.db),
		userHandler: NewUserHandler(i.db),
	}
}

type appHandler struct {
	authHandler AuthHandler
	userHandler UserHandler
}

type AppHandler interface {
	//authHandler
	CreateUser() http.HandlerFunc
	//userHandler
	GetUser() http.HandlerFunc
	UpdateUser() http.HandlerFunc
}

func (ah *appHandler) CreateUser() http.HandlerFunc {
	return ah.authHandler.CreateUser
}

func (ah *appHandler) GetUser() http.HandlerFunc {
	return ah.userHandler.GetUser
}

func (ah *appHandler) UpdateUser() http.HandlerFunc {
	return ah.userHandler.UpdateUser
}
