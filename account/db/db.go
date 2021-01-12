package db

import (
	"fmt"

	"somersbmatthews/instacook-micro/account/handler"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const database = "instacook.db"

func init() {
	db, err := gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		errMsg := fmt.Sprintf("failed to connect to database %s: %v", database, err)
		panic(errMsg)
	}

	db.AutoMigrate(&handler.Account{})

}

func Create(acc &handler.Account) err {
	
	db.Create(acc)


}
