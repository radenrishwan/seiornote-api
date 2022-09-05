package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"seiornote/common"
	"seiornote/model/web"
	"seiornote/service"
)

type UserHandler interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type userHandler struct {
	service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{UserService: userService}
}

func (handler *userHandler) Register(ctx *fiber.Ctx) error {
	var request web.RegisterUserRequest

	ctx.BodyParser(&request)

	result := handler.UserService.Register(ctx.Context(), request)

	return ctx.Status(http.StatusCreated).JSON(common.Response[web.UserSessionResponse]{
		Code:    http.StatusCreated,
		Message: "User registered successfully",
		Data:    result,
	})
}

func (handler *userHandler) Login(ctx *fiber.Ctx) error {
	var request web.LoginUserRequest

	ctx.BodyParser(&request)

	result := handler.UserService.Login(ctx.Context(), request)

	return ctx.Status(http.StatusOK).JSON(common.Response[web.UserSessionResponse]{
		Code:    http.StatusOK,
		Message: "User logged in successfully",
		Data:    result,
	})
}

func (handler *userHandler) Delete(ctx *fiber.Ctx) error {
	var request web.DeleteUserRequest

	ctx.BodyParser(&request)

	result := handler.UserService.Delete(ctx.Context(), request)

	return ctx.Status(http.StatusOK).JSON(common.Response[web.UserResponse]{
		Code:    http.StatusOK,
		Message: "User delete successfully",
		Data:    result,
	})
}

func (handler *userHandler) Update(ctx *fiber.Ctx) error {
	var request web.UpdateUserRequest

	ctx.BodyParser(&request)

	result := handler.UserService.Update(ctx.Context(), request)

	return ctx.Status(http.StatusOK).JSON(common.Response[web.UserResponse]{
		Code:    http.StatusOK,
		Message: "User update successfully",
		Data:    result,
	})
}
