package main

import (
	"TODO"
	"TODO/pkg/handler"
	"TODO/pkg/repository"
	"TODO/pkg/service"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: &s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "password",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(TODO.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running the server: &s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
