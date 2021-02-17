package handler

import (
	"context" // https://play.golang.org/p/y3wOa9O3xGr
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DATABASE = "host=localhost user=gorm password=gorm dbname=instacook port=5432 sslmode=disable"

type AccountSchema struct {
	gorm.Model
	UserID         string
	Email          string
	HashedPassword string
}

func newSchema() AccountSchema {
	return AccountSchema{}
}

func (sch AccountSchema) Create(ctx context.Context, acc *Account) error {
	db, err := gorm.Open(postgres.Open(DATABASE), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	hashedPassword := hashAndSalt(acc.password)
	ok := comparePasswords(hashedPassword, acc.password)
	if !ok {
		errMsg := fmt.Sprintf("the hashed password %s could not be decrypted", hashedPassword)
		return errors.New(errMsg)
	}

	fmt.Println(acc.email)
	fmt.Println(hashedPassword)
	schema := AccountSchema{Email: acc.email, HashedPassword: hashedPassword}
	fmt.Println(schema)
	fmt.Println(&schema)
	fmt.Println("THIS IS SPEW")
	spew.Dump(&schema)
	result := db.Create(&schema)

	fmt.Println("THIS IS RESULT")
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	return nil
}

func (sch AccountSchema) Find(ctx context.Context, acc *Account) error {
	db, err := gorm.Open(postgres.Open(DATABASE), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	hashedPassword := hashAndSalt(acc.password)
	ok := comparePasswords(hashedPassword, acc.password)
	if !ok {
		errMsg := fmt.Sprintf("the hashed password %s could not be decrypted", hashedPassword)
		return errors.New(errMsg)
	}
	sch.Email = acc.email
	sch.HashedPassword = hashedPassword
	result := db.Where(&sch).First(&sch)
	fmt.Println("THIS IS RESULT FROM AccountSchema.Find")
	fmt.Println(result)
	return nil
}

func (sch AccountSchema) Update(ctx context.Context, acc *Account, newAcc *Account) error {
	db, err := gorm.Open(postgres.Open(DATABASE), &gorm.Config{})
	if err != nil {
		return err
	}
	hashedPassword := hashAndSalt(acc.password)
	ok := comparePasswords(hashedPassword, acc.password)
	if !ok {
		errMsg := fmt.Sprintf("the hashed password %s could not be decrypted", hashedPassword)
		return errors.New(errMsg)
	}
	sch.Email = acc.email
	sch.HashedPassword = hashedPassword

	newAccSchema := newSchema()
	hashedPassword = hashAndSalt(newAcc.password)
	ok = comparePasswords(hashedPassword, newAcc.password)
	if !ok {
		errMsg := fmt.Sprintf("the hashed password %s could not be decrypted", hashedPassword)
		return errors.New(errMsg)
	}
	newAccSchema.Email = newAcc.email
	newAccSchema.HashedPassword = hashedPassword

	switch {
	case newAccSchema.Email != sch.Email && newAccSchema.HashedPassword != sch.HashedPassword:
		err = errors.New("cannot update both email and password at same time")
		return err
	case newAccSchema.Email != sch.Email:
		tx := db.Model(&sch).Where(&sch).Update("email", &newAccSchema.Email)
		fmt.Println("THIS IS TX")
		fmt.Println(tx)
	case newAccSchema.HashedPassword != sch.HashedPassword:
		tx := db.Model(&sch).Where(&sch).Update("hashedPassword", &newAccSchema)
		fmt.Println("THIS IS TX")
		fmt.Println(tx)
	default:
		err = errors.New("both email and password are the same as the originals")
		return err
	}

	return nil
}
