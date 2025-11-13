package config

import (
	"fmt"
	"freecreate/pgModels"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigPG() *gorm.DB {
	host, user, pwd, db, ssl, port := os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DB"), os.Getenv("PG_SSL"), os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s port=%s", host, user, pwd, db, ssl, port)

	gormPGClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting gorm to postgres")
		return nil
	}
	fmt.Println("gorm connect to pg successful!")

	mErr := gormPGClient.AutoMigrate(&pgModels.User{}, &pgModels.Creator{}, &pgModels.Writing{}, &pgModels.Content{})
	if mErr != nil {
		log.Fatal(mErr.Error())
		return nil
	}
	fmt.Println("all models successfully migrated!")

	return gormPGClient
}
