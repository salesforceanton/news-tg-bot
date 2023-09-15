package repository

import (
	"context"
	"fmt"
	"salesforceanton/news-tg-bot/pkg/model"

	"github.com/jmoiron/sqlx"
)

type SubscriptionsPostgres struct {
	db *sqlx.DB
}

func NewSubscriptionsPostgres(db *sqlx.DB) *SubscriptionsPostgres {
	return &SubscriptionsPostgres{db: db}
}

func (r *SubscriptionsPostgres) Add(ctx context.Context, userId int, source model.Source) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var sourceId int
	createSourceQuery := fmt.Sprintf(`
			INSERT INTO %s (name, feed_url) values ($1, $2) 
			RETURNING id
			ON CONFLICT (feed_url)
			DO NOTHING
		`, SOURCES_TABLE)

	row := tx.QueryRowContext(ctx, createSourceQuery, source.Name, source.FeedUrl)
	err = row.Scan(&sourceId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createSubscriptionQuery := fmt.Sprintf("INSERT INTO %s (user_id, source_id) VALUES ($1, $2)", SUBSCRIPTIONS_TABLE)
	_, err = tx.Exec(createSubscriptionQuery, userId, sourceId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return sourceId, tx.Commit()
}

func (r *SubscriptionsPostgres) GetUserSubscriptions(ctx context.Context, chatId int) ([]model.Source, error) {
	var result []model.Source

	query := fmt.Sprintf(`
			SELECT source.id, source.name, source.feed_url
			FROM %s source
			INNER JOIN %s subscr on subscr.source_id = source.id
			INNER JOIN %s user on user.id = subscr.user_id
			WHERE user.chat_id = $1
		`, SOURCES_TABLE, SUBSCRIPTIONS_TABLE, USERS_TABLE)
	if err := r.db.SelectContext(ctx, &result, query, chatId); err != nil {
		return nil, err
	}

	return result, nil
}
