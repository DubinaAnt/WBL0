package app

import (
	"WBL0/internal/cache"
	"WBL0/internal/handler"
	"WBL0/internal/nats"
	"WBL0/internal/repository"
	"WBL0/internal/service"
	"WBL0/pkg/db"
	"log"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
)

func Run() {

	if err := initConfig(); err != nil {
		log.Fatalf("init config error: %s", err.Error())
	}

	dataBase, err := db.NewPostgresDB(db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("init db error: %s", err.Error())
	}

	repos := repository.NewRepository(dataBase)
	cache := cache.NewCache(5*time.Minute, 10*time.Minute)
	service := service.NewService(repos, cache)
	handler := handler.NewHandler(service)

	natsConnect, err := nats.NewConnection()
	if err != nil {
		log.Fatalf("nats connection error: %s", err.Error())
	}
	defer func(newNatsConnect stan.Conn) {
		err = newNatsConnect.Close()
		if err != nil {
			log.Printf("close nats error: %s", err.Error())
		}
	}(natsConnect)
	nats.NewNatsSubscriber(natsConnect, repos, cache)

	server := new(Server)
	err = server.Run(viper.GetString("httpserver.port"), handler.InitRoutes())
	if err != nil {
		log.Fatalf("start http server error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
