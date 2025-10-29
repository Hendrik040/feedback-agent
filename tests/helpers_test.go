package tests

import (
	"encoding/json"
	"testing"
)

// assertJSONEqual compares two JSON strings for semantic equality
func assertJSONEqual(t *testing.T, expected, actual string) {
	t.Helper()

	var expectedJSON, actualJSON interface{}

	if err := json.Unmarshal([]byte(expected), &expectedJSON); err != nil {
		t.Fatalf("failed to unmarshal expected JSON: %v", err)
	}

	if err := json.Unmarshal([]byte(actual), &actualJSON); err != nil {
		t.Fatalf("failed to unmarshal actual JSON: %v", err)
	}

	expectedBytes, _ := json.Marshal(expectedJSON)
	actualBytes, _ := json.Marshal(actualJSON)

	if string(expectedBytes) != string(actualBytes) {
		t.Errorf("\nExpected JSON:\n%s\n\nActual JSON:\n%s", expected, actual)
	}
}

// floatPtr returns a pointer to a float64
func floatPtr(f float64) *float64 {
	return &f
}

// intPtr returns a pointer to an int
func intPtr(i int) *int {
	return &i
}

// stringPtr returns a pointer to a string
func stringPtr(s string) *string {
	return &s
}

