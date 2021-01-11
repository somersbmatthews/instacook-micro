package main

import (
	"nlp/handler"
	pb "nlp/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("nlp"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterNlpHandler(srv.Server(), new(handler.Nlp))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
