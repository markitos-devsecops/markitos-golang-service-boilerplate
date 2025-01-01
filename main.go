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

	log.Println("['.']:>--- </starting markitos-golang-service-boilerplate>")
	log.Println("['.']:>--------------------------------------------")
	log.Println("['.']:>")
	server := loadServer(config, repository)
	err = server.Run()
	if err != nil {
		log.Fatal("unable to start server: ", err)
	}
}

func loadServer(config configuration.MarkitosGolangServiceBoilerplateConfig, repository *database.BoilerPostgresRepository) *api.Server {
	log.Println("['.']:>----- <server.api>")
	gin.SetMode(gin.ReleaseMode)
	server := api.NewServer(config.AppAddress, repository)
	log.Println("['.']:>------- New server created")
	log.Println("['.']:>----- </server.api>")
	return server
}

func loadDatabase(config configuration.MarkitosGolangServiceBoilerplateConfig) (*database.BoilerPostgresRepository, error) {
	log.Println("['.']:>")
	log.Println("['.']:>----- <database>")
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
	log.Println("['.']:>----- </database>")
	log.Println("['.']:>")

	return repository, nil
}

func loadConfiguration() configuration.MarkitosGolangServiceBoilerplateConfig {
	log.Println("['.']:>")
	log.Println("['.']:>----- <configuration>")
	config, err := configuration.LoadConfiguration(".")
	if err != nil {
		log.Fatal("['.']:>------- unable to load configuration: ", err)
	}
	log.Println("['.']:>------- all values ready to use :)")
	log.Println("['.']:>------- serverAddress: ", config.AppAddress)
	log.Println("['.']:>----- </configuration>")
	log.Println("['.']:>")

	return config
}
