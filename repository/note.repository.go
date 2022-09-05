package repository

import (
	"context"
	"database/sql"
	"errors"
	"seiornote/helper"
	"seiornote/model/domain"
)

type NoteRepository interface {
	Save(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note
	Update(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note
	Delete(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note
	FindById(ctx context.Context, tx *sql.Tx, note domain.Note) (domain.Note, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, note domain.Note, limit int) []domain.Note
	FindFavoriteByUserId(ctx context.Context, tx *sql.Tx, note domain.Note, limit int) []domain.Note
}

type noteRepository struct {
}

func NewNoteRepository() NoteRepository {
	return &noteRepository{}
}

func (repository *noteRepository) Save(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note {
	query := "insert into notes values ($1, $2, $3, $4, $5, $6, $7)"

	_, err := tx.ExecContext(ctx, query, note.Id, note.UserId, note.Title, note.Content, note.Favorite, note.CreatedAt, note.UpdatedAt)
	helper.PanicIfError(err)

	return note
}

func (repository *noteRepository) Update(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note {
	query := "update notes set title = $1, content = $2, favorite = $3, updated_at = $4 where id = $5"

	_, err := tx.ExecContext(ctx, query, note.Title, note.Content, note.Favorite, note.UpdatedAt, note.Id)
	helper.PanicIfError(err)

	return note
}

func (repository *noteRepository) Delete(ctx context.Context, tx *sql.Tx, note domain.Note) domain.Note {
	query := "delete from notes where id = $1"

	_, err := tx.ExecContext(ctx, query, note.Id)
	helper.PanicIfError(err)

	return note
}

func (repository *noteRepository) FindById(ctx context.Context, tx *sql.Tx, note domain.Note) (domain.Note, error) {
	query := " select * from notes where id = $1"

	rows, err := tx.QueryContext(ctx, query, note.Id)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&note.Id, &note.UserId, &note.Title, &note.Content, &note.Favorite, &note.CreatedAt, &note.UpdatedAt)
		helper.PanicIfError(err)

		return note, nil
	}

	return note, errors.New("note not found")
}

func (repository *noteRepository) FindByUserId(ctx context.Context, tx *sql.Tx, note domain.Note, limit int) []domain.Note {
	query := " select * from notes where user_id = $1 limit $2"
	var notes []domain.Note

	rows, err := tx.QueryContext(ctx, query, note.UserId, limit)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		var note domain.Note
		err := rows.Scan(&note.Id, &note.UserId, &note.Title, &note.Content, &note.Favorite, &note.CreatedAt, &note.UpdatedAt)
		helper.PanicIfError(err)
		notes = append(notes, note)
	}

	return notes
}

func (repository *noteRepository) FindFavoriteByUserId(ctx context.Context, tx *sql.Tx, note domain.Note, limit int) []domain.Note {
	query := " select * from notes where user_id = $1 and favorite = true limit $2"
	var notes []domain.Note

	rows, err := tx.QueryContext(ctx, query, note.UserId, limit)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		var note domain.Note
		err := rows.Scan(&note.Id, &note.UserId, &note.Title, &note.Content, &note.Favorite, &note.CreatedAt, &note.UpdatedAt)
		helper.PanicIfError(err)
		notes = append(notes, note)
	}

	return notes
}
