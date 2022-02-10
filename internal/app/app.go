package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/alexm24/cache-img/internal/config"
	"github.com/alexm24/cache-img/internal/handler"
	"github.com/alexm24/cache-img/internal/models"
	"github.com/alexm24/cache-img/internal/repository"
	"github.com/alexm24/cache-img/internal/repository/postgres"
	"github.com/alexm24/cache-img/internal/server"
	"github.com/alexm24/cache-img/internal/service"
)

func Run(configPath string) {
	cfg, err := config.ParseConfig(configPath)

	if err != nil {
		log.Panicf("error read config: %s", err.Error())
	}

	cfgDB := models.DBConfig{
		Host:     cfg.DBConfig.Host,
		Port:     cfg.DBConfig.Port,
		DBName:   cfg.DBConfig.DBName,
		Username: cfg.DBConfig.Username,
		Password: cfg.DBConfig.Password,
		SSLMode:  cfg.DBConfig.SSLMode,
	}

	db, err := postgres.NewPostgresDB(cfgDB)

	if err != nil {
		log.Panicf("failed to initialize  db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	habdlers := handler.NewHandler(services)

	srv := new(server.Server)

	go func() {
		cfgHTTP := models.HTTPServerConfig{
			BasePath: cfg.HTTPServerConfig.BasePath,
			Port:     cfg.HTTPServerConfig.Port,
		}

		log.Printf("start http server port: %s", cfgHTTP.Port)

		err = srv.Run(cfgHTTP, habdlers.InitRoutes(cfgHTTP.BasePath))
		if err != nil {
			log.Panicf("error occured while running http server: %s:", err.Error())
		}
	}()

	signalLisner := make(chan os.Signal, 1)
	signal.Notify(signalLisner,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGHUP)
	stop := <-signalLisner
	log.Printf("stop app: %s", stop)
}
