package main

import (
	"github.com/fanfaronDo/to_do/internal/config"
	"github.com/fanfaronDo/to_do/internal/handler"
	"github.com/fanfaronDo/to_do/internal/repository"
	"github.com/fanfaronDo/to_do/internal/server"
	"github.com/fanfaronDo/to_do/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, _ := config.ConfigLoad()
	server := server.Server{}
	conn, err := repository.NewPostgres(cfg.Postgres)
	defer conn.Close()

	if err != nil {
		log.Printf("Failed to connect to postgres: %v", err)
		return
	}
	repo := repository.NewRepository(conn)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	route := handler.InitRoutes()

	go func() {
		if err = server.Run(cfg.HttpServer, route); err != nil {
			log.Printf("Failed to start server: %v", err)
			return
		}
	}()

	defer server.Shutdown(nil)
	log.Printf("Server started on %s\n", "http://"+cfg.HttpServer.Address+":"+cfg.HttpServer.Port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
}
