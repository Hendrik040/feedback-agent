package enum

import "fmt"

type Sentiment int

const (
	SentimentNeutral  Sentiment = 0
	SentimentPositive Sentiment = 1
	SentimentNegative Sentiment = 2
)

var sentimentIDs = map[Sentiment]string{
	SentimentNeutral:  "neutral",
	SentimentPositive: "positive",
	SentimentNegative: "negative",
}

var sentimentNames = map[string]Sentiment{
	"neutral":  SentimentNeutral,
	"positive": SentimentPositive,
	"negative": SentimentNegative,
}

func (s Sentiment) MarshalText() ([]byte, error) {
	if id, ok := sentimentIDs[s]; ok {
		return []byte(id), nil
	}
	return nil, fmt.Errorf("unknown Sentiment: %d", s)
}

func (s *Sentiment) UnmarshalText(text []byte) error {
	if v, ok := sentimentNames[string(text)]; ok {
		*s = v
		return nil
	}
	return fmt.Errorf("unknown Sentiment: %s", string(text))
}

func (s Sentiment) Name() string {
	if id, ok := sentimentIDs[s]; ok {
		return id
	}
	return "unknown"
}

func (s Sentiment) Category() string {
	switch s {
	case SentimentPositive:
		return "Kudos & Carrots"
	case SentimentNegative:
		return "Critiques"
	default:
		return "Context"
	}
}


