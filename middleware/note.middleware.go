package middleware

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"seiornote/common"
	"seiornote/database"
	"seiornote/exception"
	"seiornote/helper"
	"seiornote/repository"
	"strings"
)

func NoteMiddleware(ctx *fiber.Ctx) error {
	// TODO: re-write later
	db := database.GetDatabase()
	sessionRepository := repository.NewSessionRepository()

	tx, err := db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	dummy := ctx.Get("Authorization", "")

	if dummy == "" {
		return ctx.JSON(common.Response[string]{
			Code:    http.StatusUnauthorized,
			Data:    "Unauthorize",
			Message: "Token cannot be null or wrong, please check again",
		})
	}

	token := strings.Split(dummy, " ")

	session, err := helper.ClaimToken(token[1])
	if err != nil {
		panic(exception.NewValidateException(err.Error()))
	}

	result, err := sessionRepository.FindById(ctx.Context(), tx, session)
	if err != nil {
		panic(exception.NewUserException(err.Error()))
	}

	if result.ExpiredAt < helper.GenerateMilisTimeNow() {
		panic(exception.NewUserException("Token Expired, please login again"))
	}

	return ctx.Next()
}
