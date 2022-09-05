package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"seiornote/common"
	"seiornote/exception"
	"seiornote/helper"
	"seiornote/model/web"
	"seiornote/service"
	"strconv"
)

type NoteHandler interface {
	Create(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	FindByUserId(ctx *fiber.Ctx) error
	FindFavorite(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type noteHandler struct {
	service.NoteService
}

func NewNoteHandler(noteService service.NoteService) NoteHandler {
	return &noteHandler{NoteService: noteService}
}

func (handler *noteHandler) Create(ctx *fiber.Ctx) error {
	var request web.CreateNoteRequest

	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	result := handler.NoteService.Create(ctx.Context(), request, helper.GetToken(ctx))

	return ctx.Status(http.StatusCreated).JSON(common.Response[web.NoteResponse]{
		Code:    http.StatusCreated,
		Message: "Note created",
		Data:    result,
	})
}

func (handler *noteHandler) FindById(ctx *fiber.Ctx) error {
	var request web.FindByIdNoteRequest

	id := ctx.Query("id")

	request.Id = id

	result := handler.NoteService.FindById(ctx.Context(), request, helper.GetToken(ctx))

	return ctx.Status(http.StatusOK).JSON(common.Response[web.NoteResponse]{
		Code:    http.StatusOK,
		Message: "Note found",
		Data:    result,
	})
}

func (handler *noteHandler) FindByUserId(ctx *fiber.Ctx) error {
	var request web.FindByUserIdNoteRequest

	userId := ctx.Query("user_id")
	limit := ctx.Query("limit", "10")

	l, err := strconv.Atoi(limit)
	if err != nil {
		panic(exception.NewValidateException("Limit must be number"))
	}

	request.UserId = userId
	request.Limit = l

	result := handler.NoteService.FindByUserId(ctx.Context(), request, helper.GetToken(ctx))

	return ctx.Status(http.StatusOK).JSON(common.Response[[]web.NoteResponse]{
		Code:    http.StatusOK,
		Message: "Note found",
		Data:    result,
	})
}

func (handler *noteHandler) FindFavorite(ctx *fiber.Ctx) error {
	var request web.FindFavoriteNoteRequest

	userId := ctx.Query("user_id")
	limit := ctx.Query("limit", "10")

	l, err := strconv.Atoi(limit)
	if err != nil {
		panic(exception.NewValidateException("Limit must be number"))
	}

	request.UserId = userId
	request.Limit = l

	result := handler.NoteService.FindFavoriteByUserId(ctx.Context(), request, helper.GetToken(ctx))

	return ctx.Status(http.StatusOK).JSON(common.Response[[]web.NoteResponse]{
		Code:    http.StatusOK,
		Message: "Note found",
		Data:    result,
	})
}

func (handler *noteHandler) Update(ctx *fiber.Ctx) error {
	var request web.UpdateNoteRequest

	err := ctx.BodyParser(&request)
	helper.PanicIfError(err)

	result := handler.NoteService.Update(ctx.Context(), request, helper.GetToken(ctx))

	return ctx.Status(http.StatusCreated).JSON(common.Response[web.NoteResponse]{
		Code:    http.StatusCreated,
		Message: "Note created",
		Data:    result,
	})
}

func (handler *noteHandler) Delete(ctx *fiber.Ctx) error {
	var request web.DeleteNoteRequest

	id := ctx.Query("id")

	request.Id = id

	result := handler.NoteService.Delete(ctx.Context(), request, helper.GetToken(ctx))

	return ctx.Status(http.StatusOK).JSON(common.Response[string]{
		Code:    http.StatusOK,
		Message: "Note found",
		Data:    result,
	})
}
