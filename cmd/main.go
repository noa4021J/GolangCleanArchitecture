package main

import (
	"CleanArchitecture_SampleApp/config"
	"CleanArchitecture_SampleApp/infrastructure/datastore"
	"CleanArchitecture_SampleApp/infrastructure/router"
	"CleanArchitecture_SampleApp/infrastructure/server"
	"CleanArchitecture_SampleApp/interface/controllers"
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
	interactor := controllers.NewInteractor(connectedDB)
	// AppHandlerの取得
	appController := interactor.NewAppController()
	// Routerの起動
	serv := server.New()
	router.BootRouter(serv, appController)
	// DBのClose
	defer func() {
		if err := connectedDB.DB.Close(); err != nil {
			log.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()
	// Server Start
	serv.Start(config.Conf.Server.Address)
}
