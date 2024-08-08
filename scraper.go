package main

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/mori-dv/RSS/internal/database"
	"log"
	"strings"
	"sync"
	"time"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping starting on %v goroutines every %s duration \n", concurrency, timeBetweenRequest)
	getUpdatesTelegram()
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Error fetching feeds from DB:", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go ScrapeFeed(wg, db, feed)
		}
		wg.Wait()
	}

}
func ScrapeFeed(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error fetching feed from DB:", err)
		return
	}
	rssFeed, err := urlToFeed(feed.Slug)
	if err != nil {
		log.Println("Error converting feed to RSS feed:", err)
		return
	}
	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		pubAt, er := time.Parse(time.RFC1123, item.PubDate)
		if er != nil {
			log.Println("Error converting feed to RSS feed:", er)
			continue
		}
		_, pErr := db.CreatePost(context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Description: description,
				PublishedAt: pubAt,
				Url:         item.Link,
				FeedID:      feed.ID,
			})
		if pErr != nil {
			if strings.Contains(pErr.Error(), "duplicate key value") {
				continue
			}
			log.Println("Error inserting feed to DB:", pErr)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
