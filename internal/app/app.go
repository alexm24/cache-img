package app

import (
	"github.com/alexm24/cache-img/internal/handler"
	"github.com/alexm24/cache-img/internal/models"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexm24/cache-img/internal/config"
	"github.com/alexm24/cache-img/internal/server"
)

func Run(configPath string) {
	cfg, err := config.ParseConfig(configPath)

	cfgHTTP := models.HTTPServerConfig{
		BasePath: cfg.HTTPServer.BasePath,
		Port:     cfg.HTTPServer.Port,
	}

	if err != nil {
		log.Panic("error read config")
	}

	habdlers := handler.NewHandler()

	srv := new(server.Server)

	go func() {
		log.Println("http server start")
		err = srv.Run(cfgHTTP, habdlers.InitRoutes())
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
