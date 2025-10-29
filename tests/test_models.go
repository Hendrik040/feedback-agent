package main

import (
	"encoding/json"
	"fmt"
	"time"

	"feedback-collector-agent/app/models/entity"
	"feedback-collector-agent/app/models/enum"
)

func main() {
	// Create a feedback instance
	score := 0.85
	feedback := entity.Feedback{
		ID:             1,
		Title:          "Add dark mode",
		Description:    "Would love to have a dark mode option",
		Type:           enum.FeedbackFeature,
		Tags:           []string{"ui", "enhancement"},
		Sentiment:      enum.SentimentPositive,
		SentimentScore: &score,
		Votes:          42,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	// Test marshaling to JSON
	jsonData, err := json.MarshalIndent(feedback, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}

	fmt.Println("✅ Created Feedback struct successfully!")
	fmt.Println("\nJSON Output:")
	fmt.Println(string(jsonData))

	// Test enum methods
	fmt.Printf("\nFeedback Type: %s\n", feedback.Type.Name())
	fmt.Printf("Sentiment: %s\n", feedback.Sentiment.Name())
	fmt.Printf("Sentiment Category: %s\n", feedback.Sentiment.Category())
}
