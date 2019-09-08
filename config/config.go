package config

import (
	"flag"
	"os"
)

type Config struct {
	Database
	Server
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
}
type Server struct {
	Address string
}

var Conf Config

func LoadConfig() {

	var addr string
	// サーバー接続設定
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()

	//　DB接続設定
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")

	Conf = Config{
		Database: Database{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			Db:       database,
		},
		Server: Server{
			Address: addr,
		},
	}

}
