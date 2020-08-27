package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"

	"gitlab.com/bimoyong/go-user/config"
	"gitlab.com/bimoyong/go-user/handler"
	"gitlab.com/bimoyong/go-user/subscriber"
)

func main() {
	service := micro.NewService()

	client.DefaultClient = service.Client()
	server.DefaultServer = service.Server()

	service.Init(
		micro.BeforeStart(config.Init),
		micro.BeforeStop(subscriber.Close),
	)

	micro.RegisterHandler(server.DefaultServer, new(handler.User))

	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
