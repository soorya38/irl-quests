// Quests are real life short missions, challenges and dares. To access each quest, certain
// number of points are required. To gather points a player must complete quests. After completing
// each quest the player is awarded the points based on the completed quest's level.

// Once a player surpasses the point threshold, they can create quests which will be reviewed and
// and added to the quest pool for others to access.

package entity

import (
	"errors"
	"time"
)

// Quests are real life short missions, challenges and dares.
type Quest struct {
	Id          int64
	Title       string
	Description string
	Level       int
	Tags        []string

	CreatedBy int64
	UpdatedBy int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewQuest creates a new quest instance.
func NewQuest(
	title string,
	description string,
	level int,
	tags []string,
) *Quest {
	return &Quest{
		Title:       title,
		Description: description,
		Level:       level,
		Tags:        tags,
	}
}

// Validate validates the quest instance.
func (q *Quest) Validate() error {
	for _, tag := range q.Tags {
		if tag == "" {
			return errors.New("tag cannot be empty")
		}
		if len(tag) > 10 {
			return errors.New("tags cannot be longer than 10 characters")
		}
	}

	switch {
	case q.Title == "":
		return errors.New("title cannot be empty")
	case len(q.Title) > 50:
		return errors.New("title cannot having more than 50 characters")
	case q.Description == "":
		return errors.New("description cannot be empty")
	case len(q.Description) > 255:
		return errors.New("description cannot having more than 255 characters")
	}
	return nil
}
