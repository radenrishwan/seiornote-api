package validation

import (
	"fmt"
	"seiornote/exception"
	"seiornote/model/web"

	validation "github.com/go-ozzo/ozzo-validation"
)

func NewCreateNoteValidation(request web.CreateNoteRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Title,
			validation.Required,
			validation.Length(1, 25).Error(fmt.Sprintf(length, 1, 25)),
		), // TODO: check if favorite is true or false
	)

	if err != nil {
		panic(exception.NewValidateException(err.Error()))
	}
}

func FindByIdNoteValidation(request web.FindByIdNoteRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Id,
			validation.Required,
			validation.NotNil,
		),
	)

	if err != nil {
		panic(exception.NewValidateException(err.Error()))
	}
}

func FindByUserIdNoteValidation(request web.FindByUserIdNoteRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.UserId,
			validation.Required,
			validation.NotNil,
		),
		validation.Field(&request.Limit,
			validation.Required,
			validation.NotNil,
		),
	)

	if err != nil {
		panic(exception.NewValidateException(err.Error()))
	}
}

func DeleteNoteValidation(request web.DeleteNoteRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Id,
			validation.Required,
			validation.NotNil,
		),
	)

	if err != nil {
		panic(exception.NewValidateException(err.Error()))
	}
}

func NewUpdateNoteValidation(request web.UpdateNoteRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Title,
			validation.Required,
			validation.Length(1, 25).Error(fmt.Sprintf(length, 1, 25)),
		),
		validation.Field(&request.Content, validation.Required),
		validation.Field(&request.Favorite, validation.NotNil, validation.Required),
	)

	if err != nil {
		panic(exception.NewValidateException(err.Error()))
	}
}

func FindFavoriteNoteValidation(request web.FindFavoriteNoteRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.UserId,
			validation.Required,
			validation.NotNil,
		),
		validation.Field(&request.Limit,
			validation.Required,
			validation.NotNil,
		),
	)

	if err != nil {
		panic(exception.NewValidateException(err.Error()))
	}
}
