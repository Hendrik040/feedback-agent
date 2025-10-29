package tests

import (
	"encoding/json"
	"testing"
	"time"

	"feedback-collector-agent/models/entity"
	"feedback-collector-agent/models/enum"
)

func TestFeedback_JSONMarshaling(t *testing.T) {
	score := 0.85
	now := time.Now()

	feedback := entity.Feedback{
		ID:             1,
		Title:          "Add dark mode",
		Description:    "Would love to have a dark mode option",
		Type:           enum.FeedbackFeature,
		Tags:           []string{"ui", "enhancement"},
		Sentiment:      enum.SentimentPositive,
		SentimentScore: &score,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(feedback)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	// Unmarshal back
	var result entity.Feedback
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	// Verify fields
	if result.ID != feedback.ID {
		t.Errorf("ID: got %d, want %d", result.ID, feedback.ID)
	}
	if result.Title != feedback.Title {
		t.Errorf("Title: got %q, want %q", result.Title, feedback.Title)
	}
	if result.Type != feedback.Type {
		t.Errorf("Type: got %v, want %v", result.Type, feedback.Type)
	}
	if result.Sentiment != feedback.Sentiment {
		t.Errorf("Sentiment: got %v, want %v", result.Sentiment, feedback.Sentiment)
	}
	if len(result.Tags) != len(feedback.Tags) {
		t.Errorf("Tags length: got %d, want %d", len(result.Tags), len(feedback.Tags))
	}
}

func TestFeedback_JSONMarshaling_WithoutOptionalFields(t *testing.T) {
	now := time.Now()

	feedback := entity.Feedback{
		ID:          1,
		Title:       "Simple feedback",
		Description: "",
		Type:        enum.FeedbackBug,
		Tags:        []string{},
		Sentiment:   enum.SentimentNeutral,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	jsonData, err := json.Marshal(feedback)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	// Check that sentimentScore is omitted
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonData, &jsonMap)
	if err != nil {
		t.Fatalf("failed to unmarshal to map: %v", err)
	}

	if _, exists := jsonMap["sentimentScore"]; exists {
		t.Errorf("sentimentScore should be omitted when nil")
	}
}

func TestFeedback_EmptyTags(t *testing.T) {
	feedback := entity.Feedback{
		Tags: []string{},
	}

	jsonData, err := json.Marshal(feedback)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var result entity.Feedback
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if result.Tags == nil {
		t.Error("Tags should be empty slice, not nil")
	}

	if len(result.Tags) != 0 {
		t.Errorf("Tags should be empty, got length %d", len(result.Tags))
	}
}

func TestFeedback_AllFeedbackTypes(t *testing.T) {
	tests := []struct {
		name         string
		feedbackType enum.FeedbackType
		expectedJSON string
	}{
		{
			name:         "Bug feedback",
			feedbackType: enum.FeedbackBug,
			expectedJSON: `"type":"bug"`,
		},
		{
			name:         "Feature feedback",
			feedbackType: enum.FeedbackFeature,
			expectedJSON: `"type":"feature"`,
		},
		{
			name:         "General feedback",
			feedbackType: enum.FeedbackGeneral,
			expectedJSON: `"type":"general"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			feedback := entity.Feedback{
				ID:        1,
				Title:     "Test",
				Type:      tt.feedbackType,
				Tags:      []string{},
				Sentiment: enum.SentimentNeutral,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			jsonData, err := json.Marshal(feedback)
			if err != nil {
				t.Fatalf("failed to marshal: %v", err)
			}

			jsonStr := string(jsonData)
			if !contains(jsonStr, tt.expectedJSON) {
				t.Errorf("expected JSON to contain %q, got %q", tt.expectedJSON, jsonStr)
			}
		})
	}
}

func TestFeedback_AllSentiments(t *testing.T) {
	tests := []struct {
		name         string
		sentiment    enum.Sentiment
		expectedJSON string
	}{
		{
			name:         "Neutral sentiment",
			sentiment:    enum.SentimentNeutral,
			expectedJSON: `"sentiment":"neutral"`,
		},
		{
			name:         "Positive sentiment",
			sentiment:    enum.SentimentPositive,
			expectedJSON: `"sentiment":"positive"`,
		},
		{
			name:         "Negative sentiment",
			sentiment:    enum.SentimentNegative,
			expectedJSON: `"sentiment":"negative"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			feedback := entity.Feedback{
				ID:        1,
				Title:     "Test",
				Type:      enum.FeedbackGeneral,
				Tags:      []string{},
				Sentiment: tt.sentiment,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			jsonData, err := json.Marshal(feedback)
			if err != nil {
				t.Fatalf("failed to marshal: %v", err)
			}

			jsonStr := string(jsonData)
			if !contains(jsonStr, tt.expectedJSON) {
				t.Errorf("expected JSON to contain %q, got %q", tt.expectedJSON, jsonStr)
			}
		})
	}
}

func TestVote_JSONMarshaling(t *testing.T) {
	userID := 42
	ipAddress := "192.168.1.1"
	now := time.Now()

	vote := entity.Vote{
		ID:         1,
		FeedbackID: 10,
		UserID:     &userID,
		IPAddress:  &ipAddress,
		CreatedAt:  now,
	}

	jsonData, err := json.Marshal(vote)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var result entity.Vote
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if result.ID != vote.ID {
		t.Errorf("ID: got %d, want %d", result.ID, vote.ID)
	}
	if result.FeedbackID != vote.FeedbackID {
		t.Errorf("FeedbackID: got %d, want %d", result.FeedbackID, vote.FeedbackID)
	}
	if result.UserID == nil || *result.UserID != *vote.UserID {
		t.Errorf("UserID mismatch")
	}
	if result.IPAddress == nil || *result.IPAddress != *vote.IPAddress {
		t.Errorf("IPAddress mismatch")
	}
}

func TestVote_JSONMarshaling_Anonymous(t *testing.T) {
	now := time.Now()

	vote := entity.Vote{
		ID:         1,
		FeedbackID: 10,
		UserID:     nil,
		IPAddress:  nil,
		CreatedAt:  now,
	}

	jsonData, err := json.Marshal(vote)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonData, &jsonMap)
	if err != nil {
		t.Fatalf("failed to unmarshal to map: %v", err)
	}

	// Optional fields should be omitted
	if _, exists := jsonMap["userId"]; exists {
		t.Errorf("userId should be omitted when nil")
	}
	if _, exists := jsonMap["ipAddress"]; exists {
		t.Errorf("ipAddress should be omitted when nil")
	}
}

func TestVote_WithUserID(t *testing.T) {
	userID := 100
	vote := entity.Vote{
		ID:         1,
		FeedbackID: 5,
		UserID:     &userID,
		CreatedAt:  time.Now(),
	}

	jsonData, err := json.Marshal(vote)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var result entity.Vote
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if result.UserID == nil {
		t.Fatal("UserID should not be nil")
	}

	if *result.UserID != userID {
		t.Errorf("UserID: got %d, want %d", *result.UserID, userID)
	}
}

func TestVote_WithIPAddress(t *testing.T) {
	ipAddress := "10.0.0.1"
	vote := entity.Vote{
		ID:         1,
		FeedbackID: 5,
		IPAddress:  &ipAddress,
		CreatedAt:  time.Now(),
	}

	jsonData, err := json.Marshal(vote)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var result entity.Vote
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if result.IPAddress == nil {
		t.Fatal("IPAddress should not be nil")
	}

	if *result.IPAddress != ipAddress {
		t.Errorf("IPAddress: got %q, want %q", *result.IPAddress, ipAddress)
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

