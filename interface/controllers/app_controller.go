package controllers

import (
	"CleanArchitecture_SampleApp/interface/controllers/middleware"
	"CleanArchitecture_SampleApp/interface/database"
	"CleanArchitecture_SampleApp/interface/network"
)

type interactor struct {
	db database.ConnectedDB
}

type Interactor interface {
	NewAppController() AppController
}

func NewInteractor(db database.ConnectedDB) Interactor {
	return &interactor{db: db}
}

func (i *interactor) NewAppController() AppController {
	return &appController{
		middleware:     middleware.NewMiddleWare(i.db),
		authController: NewAuthController(i.db),
		userController: NewUserController(i.db),
	}
}

type appController struct {
	middleware     middleware.MiddleWare
	authController AuthController
	userController UserController
}

type AppController interface {
	//authController
	CreateUser(ar network.ApiResponser)
	//userController
	GetUser(ar network.ApiResponser)
	UpdateUser(ar network.ApiResponser)
}

func (ac *appController) CreateUser(ar network.ApiResponser) {
	ac.authController.CreateUser(ar)
}

func (ac *appController) GetUser(ar network.ApiResponser) {
	ac.userController.GetUser(ac.middleware.UserAuthorize(ar))
}

func (ac *appController) UpdateUser(ar network.ApiResponser) {
	ac.userController.UpdateUser(ac.middleware.UserAuthorize(ar))
}
