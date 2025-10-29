package entity

import "time"

type Vote struct {
	ID         int        `json:"id"`
	FeedbackID int        `json:"feedbackId"`
	UserID     *int       `json:"userId,omitempty"`
	IPAddress  *string    `json:"ipAddress,omitempty"`
	CreatedAt  time.Time  `json:"createdAt"`
}


