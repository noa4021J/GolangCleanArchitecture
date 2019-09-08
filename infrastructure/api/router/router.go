package router

import (
	"CleanArchitecture_SampleApp/infrastructure/api/handler"
	"CleanArchitecture_SampleApp/infrastructure/api/middleware"
	"CleanArchitecture_SampleApp/infrastructure/server"
)

func BootRouter(s server.Server, mw middleware.MiddleWare, handler handler.AppHandler) {
	// auth
	s.Post("/auth/create", handler.CreateUser())
	// user
	s.Get("/user/get", mw.UserAuthorize(handler.GetUser()))
	s.Post("/user/update", mw.UserAuthorize(handler.UpdateUser()))
}
