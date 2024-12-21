package main

import (
	"log"
	"markitos-svc-boilerplate/internal/domain"
	"markitos-svc-boilerplate/internal/infrastructure/api"
	"markitos-svc-boilerplate/internal/infrastructure/database"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	APP_BBDD_DSN string = "host=localhost user=admin password=admin dbname=markitos-svc-boilerplate sslmode=disable TimeZone=Europe/Madrid port=5432 sslmode=disable"
	APP_ADDRESS  string = ":3000"
)

func main() {

	log.Println("['.']:>")
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>--- <starting markitos-svc-boilerplate>")
	db, err := gorm.Open(postgres.Open(APP_BBDD_DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&domain.Boiler{})
	if err != nil {
		log.Fatal(err)
	}
	repository := database.NewBoilerPostgresRepository(db)
	log.Println("['.']:>----- <database>")
	log.Println("['.']:>------- Connected to database - migrations")
	log.Println("['.']:>----- </database>")
	log.Println("['.']:>----- <server.api>")
	gin.SetMode(gin.ReleaseMode)
	server := api.NewServer(APP_ADDRESS, repository)
	log.Println("['.']:>------- New server created")
	log.Println("['.']:>----- </server.api>")
	log.Println("['.']:>--- </starting markitos-svc-boilerplate>")
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>")
	err = server.Run()
	if err != nil {
		log.Fatal("unable to start server: ", err)
	}
}
