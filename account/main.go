package main

import (
	"account/handler"
	pb "account/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("account"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterAccountHandler(srv.Server(), new(handler.Account))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
