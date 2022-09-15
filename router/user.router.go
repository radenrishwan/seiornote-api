package router

import (
	"seiornote/handler"
	"seiornote/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewUserRouter(app *fiber.App, userHandler handler.UserHandler) {
	userRouter := app.Group("/api/v1/user")

	userRouter.Post("/register", userHandler.Register)
	userRouter.Post("/login", userHandler.Login)
	userRouter.Delete("/delete", userHandler.Delete)
	userRouter.Put("/update", userHandler.Update)

	userLogoutRouter := app.Group("/api/v1/user")
	userLogoutRouter.Use(middleware.NoteMiddleware)
	userLogoutRouter.Post("/logout", userHandler.Logout)
}
