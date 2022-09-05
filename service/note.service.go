package service

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"seiornote/exception"
	"seiornote/helper"
	"seiornote/model/domain"
	"seiornote/model/web"
	"seiornote/repository"
	"seiornote/validation"
)

type NoteService interface {
	Create(ctx context.Context, request web.CreateNoteRequest, token string) web.NoteResponse
	FindById(ctx context.Context, request web.FindByIdNoteRequest, token string) web.NoteResponse
	FindByUserId(ctx context.Context, request web.FindByUserIdNoteRequest, token string) []web.NoteResponse
	FindFavoriteByUserId(ctx context.Context, request web.FindFavoriteNoteRequest, token string) []web.NoteResponse
	Delete(ctx context.Context, request web.DeleteNoteRequest, token string) string
	Update(ctx context.Context, request web.UpdateNoteRequest, token string) web.NoteResponse
}

type noteService struct {
	repository.NoteRepository
	*sql.DB
}

func NewNoteService(noteRepository repository.NoteRepository, DB *sql.DB) NoteService {
	return &noteService{NoteRepository: noteRepository, DB: DB}
}

func (service *noteService) Create(ctx context.Context, request web.CreateNoteRequest, token string) web.NoteResponse {
	validation.NewCreateNoteValidation(request)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	session, err := helper.ClaimToken(token)
	if err != nil {
		panic(exception.NewUserException(err.Error()))
	}

	note := domain.Note{
		Id:        uuid.NewString(),
		UserId:    session.UserId,
		Title:     request.Title,
		Content:   request.Content,
		Favorite:  request.Favorite,
		CreatedAt: helper.GenerateMilisTimeNow(),
		UpdatedAt: helper.GenerateMilisTimeNow(),
	}

	service.NoteRepository.Save(ctx, tx, note)

	return web.NewNoteResponse(note)
}

func (service *noteService) FindById(ctx context.Context, request web.FindByIdNoteRequest, token string) web.NoteResponse {
	validation.FindByIdNoteValidation(request)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// claim user token
	session, err := helper.ClaimToken(token)

	note, err := service.NoteRepository.FindById(ctx, tx, domain.Note{
		Id: request.Id,
	})

	if err != nil {
		panic(exception.NewNotFoundException(err.Error()))
	}

	// check id if it has an owner the note
	if session.UserId == note.UserId {
		return web.NewNoteResponse(note)
	} else {
		panic(exception.NewValidateException("you cannot have permission to get this note!"))
	}
}

func (service *noteService) FindByUserId(ctx context.Context, request web.FindByUserIdNoteRequest, token string) []web.NoteResponse {
	// TODO: pagination here
	validation.FindByUserIdNoteValidation(request)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// claim user token
	session, err := helper.ClaimToken(token)

	// check id if it has an owner the note
	if session.UserId != request.UserId {
		panic(exception.NewValidateException("you cannot have permission to get this note!"))
	}

	// find note by user id
	notes := service.NoteRepository.FindByUserId(ctx, tx, domain.Note{
		UserId: request.UserId,
	}, request.Limit)

	return web.NewNoteResponses(notes)
}

func (service *noteService) Delete(ctx context.Context, request web.DeleteNoteRequest, token string) string {
	validation.DeleteNoteValidation(request)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// claim user token
	session, err := helper.ClaimToken(token)

	note := domain.Note{
		Id: request.Id,
	}

	// find note by id
	result, err := service.NoteRepository.FindById(ctx, tx, note)
	if err != nil {
		panic(exception.NewNotFoundException(err.Error()))
	}

	// check id if it has an owner the note
	if result.UserId != session.UserId {
		panic(exception.NewValidateException("you cannot have permission to delete this note!"))
	}

	service.NoteRepository.Delete(ctx, tx, note)

	return "Success delete note"
}

func (service *noteService) Update(ctx context.Context, request web.UpdateNoteRequest, token string) web.NoteResponse {
	validation.NewUpdateNoteValidation(request)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	session, err := helper.ClaimToken(token)
	if err != nil {
		panic(exception.NewUserException(err.Error()))
	}

	note := domain.Note{Id: request.Id}

	// check if note found
	note, err = service.NoteRepository.FindById(ctx, tx, note)
	if err != nil {
		panic(exception.NewNotFoundException(err.Error()))
	}

	// check id if it has an owner the note
	if note.UserId != session.UserId {
		panic(exception.NewValidateException("you cannot have permission to update this note!"))
	}

	// update note
	note.Title = request.Title
	note.Content = request.Content
	note.Favorite = request.Favorite
	note.UpdatedAt = helper.GenerateMilisTimeNow()

	service.NoteRepository.Update(ctx, tx, note)

	return web.NewNoteResponse(note)
}

func (service *noteService) FindFavoriteByUserId(ctx context.Context, request web.FindFavoriteNoteRequest, token string) []web.NoteResponse {
	// TODO: pagination here
	validation.FindFavoriteNoteValidation(request)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	// claim user token
	session, err := helper.ClaimToken(token)

	// check id if it has an owner the note
	if session.UserId != request.UserId {
		panic(exception.NewValidateException("you cannot have permission to get this note!"))
	}

	// find note by user id
	// TODO: limit validation
	notes := service.NoteRepository.FindFavoriteByUserId(ctx, tx, domain.Note{
		UserId: request.UserId,
	}, request.Limit)

	return web.NewNoteResponses(notes)
}
