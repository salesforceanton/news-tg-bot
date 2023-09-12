package model

type User struct {
	Id     int64  `db:"id"`
	Alias  string `db:"user_alias"`
	ChatId int64  `db:"chat_id"`
}

type SubscriptionEntity struct {
	Id       int64 `db:"id"`
	UserId   int64 `db:"user_id"`
	SourceId int64 `db:"source_id"`
}

type Source struct {
	Id      int64  `db:"id"`
	Name    string `db:"name"`
	FeedUrl string `db:"feed_url"`
}

type Article struct {
	Id          int64  `db:"id"`
	SourceId    int64  `db:"name"`
	Title       string `db:"id"`
	Link        string `db:"name"`
	Summary     string `db:"feed_url"`
	PublishedAt string `db:"id"`
	CreatedAt   string `db:"name"`
	PostedAt    string `db:"feed_url"`
}
