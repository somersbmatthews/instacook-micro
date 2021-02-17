package main

import (
	"account/handler"
	pb "account/proto"
	"fmt"
	"log"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DATABASE = "host=localhost user=gorm password=gorm dbname=instacook port=5432 sslmode=disable"

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

	db, err := gorm.Open(postgres.Open(DATABASE), &gorm.Config{})
	if err != nil {
		errMsg := fmt.Sprintf("failed to connect to database %s: %v", DATABASE, err)
		panic(errMsg)
	}

	schema := handler.AccountSchema{}

	err = db.AutoMigrate(&schema)
	if err != nil {
		fmt.Println("THIS IS ERROR FOR AUTOMIGRATE")
		fmt.Println(err)
		log.Fatal(err)
	}
}
