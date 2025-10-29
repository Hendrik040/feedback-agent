package enum

import "fmt"

type FeedbackType int

const (
	FeedbackBug     FeedbackType = 0
	FeedbackFeature FeedbackType = 1
	FeedbackGeneral FeedbackType = 2
)

var feedbackTypeIDs = map[FeedbackType]string{
	FeedbackBug:     "bug",
	FeedbackFeature: "feature",
	FeedbackGeneral: "general",
}

var feedbackTypeNames = map[string]FeedbackType{
	"bug":     FeedbackBug,
	"feature": FeedbackFeature,
	"general": FeedbackGeneral,
}

func (s FeedbackType) MarshalText() ([]byte, error) {
	if id, ok := feedbackTypeIDs[s]; ok {
		return []byte(id), nil
	}
	return nil, fmt.Errorf("unknown FeedbackType: %d", s)
}

func (s *FeedbackType) UnmarshalText(text []byte) error {
	if v, ok := feedbackTypeNames[string(text)]; ok {
		*s = v
		return nil
	}
	return fmt.Errorf("unknown FeedbackType: %s", string(text))
}

func (s FeedbackType) Name() string {
	if id, ok := feedbackTypeIDs[s]; ok {
		return id
	}
	return "unknown"
}


