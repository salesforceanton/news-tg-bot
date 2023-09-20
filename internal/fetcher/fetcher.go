package fetcher

import (
	"context"
	"log"
	"salesforceanton/news-tg-bot/internal/repository"
	feed_processor "salesforceanton/news-tg-bot/internal/source"
	"salesforceanton/news-tg-bot/pkg/model"
	"sync"
	"time"

	"golang.org/x/exp/slices"
)

type Fetcher struct {
	articlesStorage repository.Articles
	sourcesStorage  repository.Sources
	interval        time.Duration
	filterKeywords  []string
	wg              *sync.WaitGroup
}

func NewFetcher(
	articlesStorage repository.Articles,
	sourcesStorage repository.Sources,
	interval time.Duration,
	filterKeywords []string) *Fetcher {
	return &Fetcher{
		articlesStorage: articlesStorage,
		sourcesStorage:  sourcesStorage,
		interval:        interval,
		filterKeywords:  filterKeywords,
		wg:              new(sync.WaitGroup),
	}
}

func (f *Fetcher) Start(ctx context.Context) error {
	ticker := time.NewTicker(f.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-ticker.C:
			if err := f.fetch(ctx); err != nil {
				return err
			}
		}
	}
}

func (f *Fetcher) fetch(ctx context.Context) error {
	sources, err := f.sourcesStorage.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, source := range sources {
		rssSource := feed_processor.NewRSSSource(source)

		f.wg.Add(1)
		// Worker which retain and store articles
		go func(source feed_processor.RSSSource) {
			defer f.wg.Done()

			feedItems, err := source.Fetch(ctx)
			if err != nil {
				log.Printf("[ERROR] failed to fetch items from source %q: %v", source.GetName(), err)
				return
			}

			if err = f.storeArticles(ctx, feedItems, source); err != nil {
				log.Printf("[ERROR] failed to save feed items from source %q: %v", source.GetName(), err)
				return
			}
		}(rssSource)
	}

	f.wg.Wait()
	return nil
}

func (f *Fetcher) storeArticles(ctx context.Context, feedItems []model.FeedItem, source feed_processor.RSSSource) error {
	for _, feedItem := range feedItems {
		if f.isShouldBeSkipped(feedItem) {
			log.Printf("[INFO] item %q (%s) from source %q should be skipped", feedItem.Title, feedItem.Link, source.GetName())
			continue
		}
		_, err := f.articlesStorage.SaveArticle(ctx, model.Article{
			SourceId:    source.GetId(),
			Title:       feedItem.Title,
			Link:        feedItem.Link,
			Summary:     feedItem.Summary,
			PublishedAt: feedItem.Date.UTC().String(),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Fetcher) isShouldBeSkipped(feedItem model.FeedItem) bool {
	for _, filter := range f.filterKeywords {
		if slices.Contains(feedItem.Categories, filter) {
			return true
		}
	}

	return false
}
