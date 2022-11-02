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

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetOutput(os.Stderr)

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

}
