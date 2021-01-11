package main

import (
	"grocery/handler"
	pb "grocery/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("grocery"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterGroceryHandler(srv.Server(), new(handler.Grocery))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
