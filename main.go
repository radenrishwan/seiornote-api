package main

import (
	"log"
	"os"
	"seiornote/database"
	"seiornote/handler"
	"seiornote/repository"
	"seiornote/router"
	"seiornote/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	db := database.GetDatabase()

	app := fiber.New(fiber.Config{
		AppName:      "seiornote-api",
		ErrorHandler: handler.ErrorHandler,
	})

	// database migration
	database.Migration()

	app.Use(logger.New())
	app.Use(recover.New())

	// TODO: add dependency injection

	// Register Repository
	userRepository := repository.NewUserRepository()
	sessionRepository := repository.NewSessionRepository()
	noteRepository := repository.NewNoteRepository()

	// Register Service
	userService := service.NewUserService(userRepository, sessionRepository, db)
	noteService := service.NewNoteService(noteRepository, db)

	// Register Handler
	userHandler := handler.NewUserHandler(userService)
	noteHandler := handler.NewNoteHandler(noteService)

	// Register routes
	router.NewUserRouter(app, userHandler)
	router.NewNoteRouter(app, noteHandler)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "App Running !!!",
			"version": "1",
		})
	})

	port := os.Getenv("APP_PORT")

	log.Fatal(app.Listen(":" + port))
}
