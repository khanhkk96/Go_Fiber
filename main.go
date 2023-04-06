package main

import (
	"fmt"
	"go-fiber-crud/config"
	"go-fiber-crud/controller"
	"go-fiber-crud/model"
	"go-fiber-crud/repository"
	"go-fiber-crud/router"
	"go-fiber-crud/service"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Print("Running service ...")
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	//Database
	db := config.ConnectDB(&loadConfig)
	validate := validator.New()
	db.Table("notes").AutoMigrate(&model.Note{})

	//Init respository
	noteRepository := repository.NewNoteRepositoryImpl(db)

	//Init service
	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	//Init controller
	noteController := controller.NewNoteController(noteService)

	//Routes
	noteRoute := router.NewRouter(noteController)

	app := fiber.New()
	app.Mount("/api", noteRoute)

	log.Fatal(app.Listen(":8000"))
}
