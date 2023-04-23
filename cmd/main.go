package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dnevsky/firstapi"
	"github.com/dnevsky/firstapi/internal/handler"
	"github.com/dnevsky/firstapi/internal/repository"
	"github.com/dnevsky/firstapi/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env vars: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to init db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(firstapi.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error while start http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Print("TodoApp shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error while stop server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error while close db connect: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
