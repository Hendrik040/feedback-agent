package entity

import (
	"time"

	"collector-agent/app/models/enum"
)

type Feedback struct {
	ID             int               `json:"id"`
	Title          string            `json:"title"`
	Description    string            `json:"description"`
	Type           enum.FeedbackType `json:"type"`
	Tags           []string          `json:"tags"`
	Sentiment      enum.Sentiment    `json:"sentiment"`
	SentimentScore *float64          `json:"sentimentScore,omitempty"`
	Votes          int               `json:"votes"`
	CreatedAt      time.Time         `json:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt"`
}

type Vote struct {
	ID         int       `json:"id"`
	FeedbackID int       `json:"feedbackId"`
	UserID     *int      `json:"userId,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
}

// GetAccumulatedVotes calculates and returns the total number of votes for this feedback
func (f *Feedback) GetAccumulatedVotes(votes []Vote) int {
	voteCount := 0
	for _, vote := range votes {
		if vote.ID == f.ID {
			voteCount++
		}
	}
	return voteCount
}
