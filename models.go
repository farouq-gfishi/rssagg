package main

import (
	"time"

	"github.com/farouq-gfishi/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_Id"`
}

type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_Id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	var feeds []Feed
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows database.FeedFollow) FeedFollows {
	return FeedFollows{
		ID:        dbFeedFollows.ID,
		CreatedAt: dbFeedFollows.CreatedAt,
		UpdatedAt: dbFeedFollows.UpdatedAt,
		UserID:    dbFeedFollows.UserID,
		FeedID:    dbFeedFollows.UserID,
	}
}
