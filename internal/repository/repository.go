package repository

import (
	"context"
	"salesforceanton/news-tg-bot/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Articles
	Subscriptions
	Users
}

type Articles interface {
	SaveArticle(ctx context.Context, article model.Article) (int, error)
	GetNotPostedArticles(ctx context.Context, sourceIds []int) ([]model.Article, error)
	MarkAsPosted(ctx context.Context, id int) error
}

type Subscriptions interface {
	Add(ctx context.Context, userId int, source model.Source) (int, error)
	GetUserSubscriptions(ctx context.Context, userId int) ([]model.Source, error)
}

type Users interface {
	CreateUser(ctx context.Context, user model.User) (int, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Articles:      NewArticlesPostgres(db),
		Subscriptions: NewSubscriptionsPostgres(db),
		Users:         NewUsersPostgres(db),
	}
}
