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
	SourceId    int64  `db:"source_id"`
	Title       string `db:"title"`
	Link        string `db:"link"`
	Summary     string `db:"summary"`
	PublishedAt string `db:"published_at"`
	CreatedAt   string `db:"created_at"`
	Posted      bool   `db:"posted"`
}
