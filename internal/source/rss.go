package feed_processor

import (
	"context"
	"salesforceanton/news-tg-bot/pkg/model"

	"github.com/SlyMarbo/rss"
	"github.com/samber/lo"
)

type RSSSource struct {
	SourceId   int64
	URL        string
	SourceName string
}

func NewRSSSource(source model.Source) RSSSource {
	return RSSSource{
		SourceId:   source.Id,
		URL:        source.FeedUrl,
		SourceName: source.Name,
	}
}

func (s *RSSSource) GetId() int64 {
	return s.SourceId
}

func (s *RSSSource) GetName() string {
	return s.SourceName
}

func (s *RSSSource) Fetch(ctx context.Context) ([]model.FeedItem, error) {
	feed, err := s.loadFeed(ctx, s.URL)
	if err != nil {
		return nil, err
	}

	return lo.Map(feed.Items, func(item *rss.Item, _ int) model.FeedItem {
		return model.FeedItem{
			Title:      item.Title,
			Categories: item.Categories,
			Link:       item.Link,
			Date:       item.Date,
			Summary:    item.Summary,
			SourceName: s.SourceName,
		}
	}), nil
}

func (s *RSSSource) loadFeed(ctx context.Context, feedUrl string) (*rss.Feed, error) {
	feedChan := make(chan *rss.Feed)
	errorChan := make(chan error)

	go func() {
		feed, err := rss.Fetch(feedUrl)
		if err != nil {
			errorChan <- err
			return
		}
		feedChan <- feed
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errorChan:
		return nil, err
	case feed := <-feedChan:
		return feed, nil
	}
}
