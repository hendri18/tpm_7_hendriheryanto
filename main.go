package main

import (
	"tpm_7_HendriHeryanto/models"
	"tpm_7_HendriHeryanto/routers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var PORT = "localhost:8080"
	// passwordnya saya sensor
	dsn := "host=localhost user=postgres password=***** dbname=tpm_6 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Product{}, &models.User{})
	if err != nil {
		panic(err)
	}

	routers.SetupRouter(db).Run(PORT)
}
