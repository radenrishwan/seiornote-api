package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"seiornote/common"
	"seiornote/exception"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(exception.NotFoundException)
	if ok {
		return ctx.Status(http.StatusNotFound).JSON(common.Response[string]{
			Code:    http.StatusNotFound,
			Message: "not found",
			Data:    err.Error(),
		})
	}

	_, ok = err.(exception.ValidateException)
	if ok {
		return ctx.Status(http.StatusBadRequest).JSON(common.Response[string]{
			Code:    http.StatusBadRequest,
			Message: "error",
			Data:    err.Error(),
		})
	}

	_, ok = err.(exception.UserException)
	if ok {
		return ctx.Status(http.StatusBadRequest).JSON(common.Response[string]{
			Code:    http.StatusBadRequest,
			Message: "error",
			Data:    err.Error(),
		})
	}

	return ctx.Status(http.StatusInternalServerError).JSON(common.Response[string]{
		Code:    http.StatusInternalServerError,
		Message: "error",
		Data:    err.Error(),
	})
}
