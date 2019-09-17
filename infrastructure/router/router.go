package router

import (
	"CleanArchitecture_SampleApp/infrastructure/server"
	"CleanArchitecture_SampleApp/interface/controllers"
)

func BootRouter(s server.Server, controller controllers.AppController) {
	// auth
	s.Post("/auth/create", func(hc *server.HttpContext) { controller.CreateUser(hc) })
	// user
	s.Get("/user/get", func(hc *server.HttpContext) { controller.GetUser(hc) })
	s.Post("/user/update", func(hc *server.HttpContext) { controller.UpdateUser(hc) })
}
