package config

import (
	"fmt"
	"freecreate/gormModels"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigGORM() *gorm.DB {
	host, user, pwd, db, ssl, port := os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DB"), os.Getenv("PG_SSL"), os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s port=%s", host, user, pwd, db, ssl, port)

	gormPGClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting gorm to postgres")
		return nil
	}
	fmt.Println("gorm connect to pg successful!")

	mErr := gormPGClient.AutoMigrate(&pgModels.User{}, &pgModels.Creator{}, &pgModels.Tag{}, &pgModels.Writing{}, &pgModels.Content{}, &pgModels.Bookshelf{}, &pgModels.ContentTag{}, &pgModels.Donation{}, &pgModels.FreecreateDonation{}, &pgModels.ReadWriting{}, &pgModels.LikedWriting{}, &pgModels.ReadingListWriting{}, &pgModels.LibraryWriting{}, &pgModels.BookshelfWriting{})
	if mErr != nil {
		log.Fatal(mErr.Error())
		return nil
	}
	fmt.Println("all models successfully migrated!")

	type Table struct {
		TableName string
	}

	var tables []Table

	tErr := gormPGClient.Table("information_schema.tables").Select("table_name").Where("table_schema = ?", "public").Find(&tables).Error
	if tErr != nil {
		log.Fatal(tErr.Error())
		return nil
	}
	fmt.Println("tables", tables)

	return gormPGClient
}
