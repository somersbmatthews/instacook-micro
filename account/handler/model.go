package handler

import (
	"context" // https://play.golang.org/p/y3wOa9O3xGr
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DATABASE = "instacook.db"

type accountSchema struct {
	gorm.Model
	email          string
	hashedPassword string
}

func init() {
	db, err := gorm.Open(postgres.Open(DATABASE), &gorm.Config{})
	if err != nil {
		errMsg := fmt.Sprintf("failed to connect to database %s: %v", DATABASE, err)
		panic(errMsg)
	}
	db.AutoMigrate(&accountSchema{})
}

func newSchema() accountSchema {
	return accountSchema{}
}

func (sch accountSchema) Create(ctx context.Context, acc *Account) error {
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
	tx := db.Create(accountSchema{email: acc.email, hashedPassword: hashedPassword})
	fmt.Println("THIS IS TX")
	fmt.Println(tx)
	return nil
}

func (sch accountSchema) Find(ctx context.Context, acc *Account) error {
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
	sch.email = acc.email
	sch.hashedPassword = hashedPassword
	result := db.Where(&sch).First(&sch)
	fmt.Println("THIS IS RESULT FROM accountSchema.Find")
	fmt.Println(result)
	return nil
}

func (sch accountSchema) Update(ctx context.Context, acc *Account, newAcc *Account) error {
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
	sch.email = acc.email
	sch.hashedPassword = hashedPassword

	newAccSchema := newSchema()
	hashedPassword = hashAndSalt(newAcc.password)
	ok = comparePasswords(hashedPassword, newAcc.password)
	if !ok {
		errMsg := fmt.Sprintf("the hashed password %s could not be decrypted", hashedPassword)
		return errors.New(errMsg)
	}
	newAccSchema.email = newAcc.email
	newAccSchema.hashedPassword = hashedPassword

	switch {
	case newAccSchema.email != sch.email && newAccSchema.hashedPassword != sch.hashedPassword:
		err = errors.New("cannot update both email and password at same time")
		return err
	case newAccSchema.email != sch.email:
		tx := db.Model(&sch).Where(&sch).Update("email", &newAccSchema.email)
		fmt.Println("THIS IS TX")
		fmt.Println(tx)
	case newAccSchema.hashedPassword != sch.hashedPassword:
		tx := db.Model(&sch).Where(&sch).Update("hashedPassword", &newAccSchema)
		fmt.Println("THIS IS TX")
		fmt.Println(tx)
	default:
		err = errors.New("both email and password are the same as the originals")
		return err
	}

	return nil
}
