package router

import (
	"github.com/gofiber/fiber/v2"
	"seiornote/handler"
	"seiornote/middleware"
)

func NewNoteRouter(app *fiber.App, noteHandler handler.NoteHandler) {
	noteRouter := app.Group("/api/v1/note")

	noteRouter.Use(middleware.NoteMiddleware)
	noteRouter.Post("/", noteHandler.Create)

	noteRouter.Get("/", noteHandler.FindById)
	noteRouter.Get("/favorite", noteHandler.FindFavorite)

	noteRouter.Put("/", noteHandler.Update)
	noteRouter.Delete("/", noteHandler.Delete)

	// find by user id
	notesRouter := app.Group("/api/v1/notes")
	notesRouter.Use(middleware.NoteMiddleware)

	notesRouter.Get("/", noteHandler.FindByUserId)
}
