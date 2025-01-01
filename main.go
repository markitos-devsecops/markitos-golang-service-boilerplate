package main

import (
	"log"
	"markitos-golang-service-boilerplate/infrastructure/api"
	"markitos-golang-service-boilerplate/infrastructure/configuration"
	"markitos-golang-service-boilerplate/infrastructure/database"
	"markitos-golang-service-boilerplate/internal/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("['.']:>")
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>--- <starting markitos-golang-service-boilerplate>")

	config := loadConfiguration()
	repository, err := loadDatabase(config)
	if err != nil {
		log.Fatal(err)
	}

	server := loadServer(config, repository)
	log.Println("['.']:>--- </starting markitos-golang-service-boilerplate>")
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>")
	err = server.Run()
	if err != nil {
		log.Fatal("unable to start server: ", err)
	}
}

func loadServer(config configuration.MarkitosGolangServiceBoilerplateConfig, repository *database.BoilerPostgresRepository) *api.Server {
	gin.SetMode(gin.ReleaseMode)
	server := api.NewServer(config.AppAddress, repository)
	log.Println("['.']:>------- New server created")
	return server
}

func loadDatabase(config configuration.MarkitosGolangServiceBoilerplateConfig) (*database.BoilerPostgresRepository, error) {
	db, err := gorm.Open(postgres.Open(config.DsnDatabase), &gorm.Config{})
	if err != nil {
		log.Fatal("['.']:> error unable to connect to database:", err)
	}
	err = db.AutoMigrate(&domain.Boiler{})
	if err != nil {
		log.Fatal("['.']:> error unable to migrate database:", err)
	}
	repository := database.NewBoilerPostgresRepository(db)
	log.Println("['.']:>------- Connected to database - migrations")

	return repository, nil
}

func loadConfiguration() configuration.MarkitosGolangServiceBoilerplateConfig {
	config, err := configuration.LoadConfiguration(".")
	if err != nil {
		log.Fatal("['.']:>------- unable to load configuration: ", err)
	}
	log.Println("['.']:>------- all values ready to use :)")
	log.Println("['.']:>------- serverAddress: ", config.AppAddress)

	return config
}
