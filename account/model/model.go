package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const database = "instacook.db"

type AccountSchema struct {
	gorm.Model
	email          string
	hashedPassword string
}

func init() {
	db, err := gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		errMsg := fmt.Sprintf("failed to connect to database %s: %v", database, err)
		panic(errMsg)
	}
	db.AutoMigrate(&AccountSchema{})
}

func Create(acc *AccountSchema) error {
	db, err := gorm.Open(postgres.Open(database), &gorm.Config{})
	db.Create(acc)
	return err

}
