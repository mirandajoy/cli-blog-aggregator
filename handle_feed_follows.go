package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mirandjoy/gator/internal/database"
)

func handleFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("not enough arguments were provided")
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("feed couldn't be followed: %w", err)
	}

	fmt.Printf("%v feed followed by %v", feed.Name, s.cfg.CurrentUserName)

	return nil
}

func handleGetFollowFeeds(s *state, cmd command, user database.User) error {
	userFeeds, err := s.db.GetFollowedFeeds(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("feeds couldn't be fetched: %w", err)
	}

	for _, feed := range userFeeds {
		fmt.Printf("%v\n", feed.Name)
	}

	return nil
}

func handleUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("not enough arguments were provided")
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	err = s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete feed: %w", err)
	}

	return nil
}
