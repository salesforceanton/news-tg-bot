package repository

import (
	"context"
	"fmt"
	"salesforceanton/news-tg-bot/pkg/model"

	"github.com/jmoiron/sqlx"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (r *UsersPostgres) CreateUser(ctx context.Context, user model.User) (int, error) {
	var result int

	query := fmt.Sprintf("INSERT INTO %s (chat_id, user_alias) VALUES ($1, $2) RETURNING id", USERS_TABLE)
	row := r.db.QueryRowContext(ctx, query, user.ChatId, user.Alias)

	if err := row.Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}
