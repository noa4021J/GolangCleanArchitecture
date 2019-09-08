package main

import (
	"CleanArchitecture_SampleApp/config"
	"CleanArchitecture_SampleApp/infrastructure/api/handler"
	"CleanArchitecture_SampleApp/infrastructure/api/middleware"
	"CleanArchitecture_SampleApp/infrastructure/api/router"
	"CleanArchitecture_SampleApp/infrastructure/datastore"
	"CleanArchitecture_SampleApp/infrastructure/server"
	"fmt"
	"log"
)

func main() {
	// DB情報・サーバー情報の読み込み
	config.LoadConfig()
	fmt.Println(config.Conf.Db)
	// DBの起動
	connectedDB := datastore.BootMysqlDB()
	// intaractorを作成
	interactor := handler.NewInteractor(connectedDB)
	// Middlewareの起動
	middleware := middleware.NewMiddleWare(connectedDB)
	// AppHandlerの取得
	appHandler := interactor.NewAppHandler()
	// Routerの起動
	serv := server.New()
	router.BootRouter(serv, middleware, appHandler)
	// DBのClose
	defer func() {
		if err := connectedDB.DB.Close(); err != nil {
			log.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()
	// Server Start
	serv.Start(config.Conf.Server.Address)
}
