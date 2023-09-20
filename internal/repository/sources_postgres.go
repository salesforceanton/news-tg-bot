package repository

import (
	"context"
	"fmt"
	"salesforceanton/news-tg-bot/pkg/model"

	"github.com/jmoiron/sqlx"
)

type SourcesPostgres struct {
	db *sqlx.DB
}

func NewSourcesPostgres(db *sqlx.DB) *SourcesPostgres {
	return &SourcesPostgres{
		db: db,
	}
}

func (r *SourcesPostgres) GetAll(ctx context.Context) ([]model.Source, error) {
	var result []model.Source

	query := fmt.Sprintf(`
			SELECT source.id, source.name, source.feed_url
			FROM %s
		`, SOURCES_TABLE)
	if err := r.db.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}

	return result, nil
}
