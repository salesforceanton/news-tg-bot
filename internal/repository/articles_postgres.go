package repository

import (
	"context"
	"fmt"
	"salesforceanton/news-tg-bot/pkg/model"

	"github.com/jmoiron/sqlx"
)

type ArticlesPostgres struct {
	db *sqlx.DB
}

func NewArticlesPostgres(db *sqlx.DB) *ArticlesPostgres {
	return &ArticlesPostgres{db: db}
}

func (r *ArticlesPostgres) SaveArticle(ctx context.Context, article model.Article) (int, error) {
	var result int

	query := fmt.Sprintf(`
		INSERT INTO %s (source_id, title, link, summary, publiched_at, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
		`, ARTICLES_TABLE)

	row := r.db.QueryRow(query, article.SourceId, article.Link, article.Title, article.Summary, article.PublishedAt, article.CreatedAt)

	if err := row.Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}

func (r *ArticlesPostgres) GetNotPostedArticles(ctx context.Context, sourceIds []int) ([]model.Article, error) {
	var result []model.Article

	query := fmt.Sprintf(`
			SELECT source_id, title, link, summary, publiched_at, created_at 
			FROM %s
			WHERE source_id IN $1
			AND posted IS NOT TRUE
		`, ARTICLES_TABLE)
	if err := r.db.SelectContext(ctx, &result, query, sourceIds); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ArticlesPostgres) MarkAsPosted(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s SET posted_at = TRUE WHERE id = $1;,
	`, ARTICLES_TABLE)

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
