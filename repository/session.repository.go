package repository

import (
	"context"
	"database/sql"
	"errors"
	"seiornote/helper"
	"seiornote/model/domain"
)

type SessionRepository interface {
	Save(ctx context.Context, tx *sql.Tx, session domain.Session) domain.Session
	FindById(ctx context.Context, tx *sql.Tx, session domain.Session) (domain.Session, error)
	Delete(ctx context.Context, tx *sql.Tx, session domain.Session) domain.Session
}

type sessionRepository struct {
}

func NewSessionRepository() SessionRepository {
	return &sessionRepository{}
}

func (repository *sessionRepository) Save(ctx context.Context, tx *sql.Tx, session domain.Session) domain.Session {
	query := "insert into sessions values ($1, $2, $3, $4)"

	_, err := tx.ExecContext(ctx, query, session.Id, session.UserId, session.CreatedAt, session.ExpiredAt)
	helper.PanicIfError(err)

	return session
}

func (repository *sessionRepository) FindById(ctx context.Context, tx *sql.Tx, session domain.Session) (domain.Session, error) {
	query := " select * from sessions where user_id = $1"

	rows, err := tx.QueryContext(ctx, query, session.UserId)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&session.Id, &session.UserId, &session.CreatedAt, &session.ExpiredAt)
		helper.PanicIfError(err)

		return session, nil
	}

	return session, errors.New("user not found")
}

func (repository *sessionRepository) Delete(ctx context.Context, tx *sql.Tx, session domain.Session) domain.Session {
	query := "delete from sessions where user_id = $1"

	_, err := tx.ExecContext(ctx, query, session.UserId)
	helper.PanicIfError(err)

	return session
}
