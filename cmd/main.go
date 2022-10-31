package main

import (
	"course_work"
	"course_work/pkg/handler"
	"course_work/pkg/repository"
	"course_work/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config : %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error env variables : %s", err.Error())
	}

	db, err := repository.NewPostgresDb()
	if err != nil {
		log.Fatalf("failed to initialize db : %s", err.Error())
	}

	repos := repository.NewRepository(db)    // working with db
	services := service.NewService(repos)    // business logic
	handlers := handler.NewHandler(services) // http request

	srv := new(course_work.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error to running http server : %s", err)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
