package tests

import (
	"testing"

	"feedback-collector-agent/models/enum"
)

func TestFeedbackType_MarshalText(t *testing.T) {
	tests := []struct {
		name     string
		input    enum.FeedbackType
		expected string
		wantErr  bool
	}{
		{
			name:     "Bug type",
			input:    enum.FeedbackBug,
			expected: "bug",
			wantErr:  false,
		},
		{
			name:     "Feature type",
			input:    enum.FeedbackFeature,
			expected: "feature",
			wantErr:  false,
		},
		{
			name:     "General type",
			input:    enum.FeedbackGeneral,
			expected: "general",
			wantErr:  false,
		},
		{
			name:     "Invalid type",
			input:    enum.FeedbackType(999),
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.MarshalText()

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if string(result) != tt.expected {
				t.Errorf("got %q, want %q", string(result), tt.expected)
			}
		})
	}
}

func TestFeedbackType_UnmarshalText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected enum.FeedbackType
		wantErr  bool
	}{
		{
			name:     "Valid bug",
			input:    "bug",
			expected: enum.FeedbackBug,
			wantErr:  false,
		},
		{
			name:     "Valid feature",
			input:    "feature",
			expected: enum.FeedbackFeature,
			wantErr:  false,
		},
		{
			name:     "Valid general",
			input:    "general",
			expected: enum.FeedbackGeneral,
			wantErr:  false,
		},
		{
			name:     "Invalid type",
			input:    "invalid",
			expected: enum.FeedbackType(0),
			wantErr:  true,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: enum.FeedbackType(0),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result enum.FeedbackType
			err := result.UnmarshalText([]byte(tt.input))

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFeedbackType_Name(t *testing.T) {
	tests := []struct {
		name     string
		input    enum.FeedbackType
		expected string
	}{
		{"Bug name", enum.FeedbackBug, "bug"},
		{"Feature name", enum.FeedbackFeature, "feature"},
		{"General name", enum.FeedbackGeneral, "general"},
		{"Invalid name", enum.FeedbackType(999), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Name()
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestSentiment_MarshalText(t *testing.T) {
	tests := []struct {
		name     string
		input    enum.Sentiment
		expected string
		wantErr  bool
	}{
		{
			name:     "Neutral sentiment",
			input:    enum.SentimentNeutral,
			expected: "neutral",
			wantErr:  false,
		},
		{
			name:     "Positive sentiment",
			input:    enum.SentimentPositive,
			expected: "positive",
			wantErr:  false,
		},
		{
			name:     "Negative sentiment",
			input:    enum.SentimentNegative,
			expected: "negative",
			wantErr:  false,
		},
		{
			name:     "Invalid sentiment",
			input:    enum.Sentiment(999),
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.MarshalText()

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if string(result) != tt.expected {
				t.Errorf("got %q, want %q", string(result), tt.expected)
			}
		})
	}
}

func TestSentiment_UnmarshalText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected enum.Sentiment
		wantErr  bool
	}{
		{
			name:     "Valid neutral",
			input:    "neutral",
			expected: enum.SentimentNeutral,
			wantErr:  false,
		},
		{
			name:     "Valid positive",
			input:    "positive",
			expected: enum.SentimentPositive,
			wantErr:  false,
		},
		{
			name:     "Valid negative",
			input:    "negative",
			expected: enum.SentimentNegative,
			wantErr:  false,
		},
		{
			name:     "Invalid sentiment",
			input:    "invalid",
			expected: enum.Sentiment(0),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result enum.Sentiment
			err := result.UnmarshalText([]byte(tt.input))

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSentiment_Category(t *testing.T) {
	tests := []struct {
		name     string
		input    enum.Sentiment
		expected string
	}{
		{"Neutral category", enum.SentimentNeutral, "Context"},
		{"Positive category", enum.SentimentPositive, "Kudos & Carrots"},
		{"Negative category", enum.SentimentNegative, "Critiques"},
		{"Invalid category defaults to Context", enum.Sentiment(999), "Context"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Category()
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestSentiment_Name(t *testing.T) {
	tests := []struct {
		name     string
		input    enum.Sentiment
		expected string
	}{
		{"Neutral name", enum.SentimentNeutral, "neutral"},
		{"Positive name", enum.SentimentPositive, "positive"},
		{"Negative name", enum.SentimentNegative, "negative"},
		{"Invalid name", enum.Sentiment(999), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Name()
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

