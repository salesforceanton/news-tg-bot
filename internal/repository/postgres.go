package repository

import (
	"fmt"
	"salesforceanton/news-tg-bot/internal/config"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const (
	POSTGRESS_DB_TYPE   = "postgres"
	ARTICLES_TABLE      = "articles"
	SOURCES_TABLE       = "sources"
	SUBSCRIPTIONS_TABLE = "subscriptions"
)

func NewPostgresDB(cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", POSTGRESS_DB_TYPE, cfg.Username, cfg.Password, cfg.Host, cfg.Name))
	db, err := sqlx.Open(POSTGRESS_DB_TYPE, pgUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
