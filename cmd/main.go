package main

import (
	"course_work"
	"course_work/configs"
	"course_work/pkg/handler"
	"course_work/pkg/repository"
	"course_work/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// @title Course Work Computer Store Back-end
// @version 1.0
// @description Api server for Computer store

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name

func main() {

	file, err := os.OpenFile("logs/log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)

	logrus.SetOutput(file)
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := configs.InitConfig(); err != nil {
		logrus.Fatalf("error initializing config : %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error env variables : %s", err.Error())
	}

	db, err := repository.NewPostgresDb()
	if err != nil {
		logrus.Fatalf("failed to initialize db : %s", err.Error())
	}

	repos := repository.NewRepository(db)    // working with db
	services := service.NewService(repos)    // business logic
	handlers := handler.NewHandler(services) // http request

	srv := new(course_work.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error to running http server : %s", err)
	}
	logrus.Infof("Server working on @localhost:%s", viper.GetString("port"))

}
