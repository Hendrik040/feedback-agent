# Test Suite Documentation

This directory contains comprehensive unit tests for the feedback collector agent models.

## Test Structure

```
tests/
├── enum_test.go      # Tests for enum types (FeedbackType, Sentiment)
├── entity_test.go    # Tests for entity structs (Feedback, Vote)
└── helpers_test.go   # Shared test utilities and helper functions
```

## Running Tests

### Using Make (Recommended)

```bash
# Run all tests with verbose output
make test

# Run tests without verbose output
make test-short

# Run tests with HTML coverage report
make test-coverage

# Run tests with text coverage summary
make test-coverage-text

# Run the demo example
make demo
```

### Using Go Directly

```bash
# Run all tests with verbose output
go test ./tests/... -v

# Run all tests
go test ./tests/...

# Run specific test
go test ./tests/... -v -run TestFeedbackType_MarshalText

# Run with coverage
go test ./tests/... -cover

# Generate coverage profile
go test ./tests/... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Test Coverage

The test suite covers:

### Enum Tests (`enum_test.go`)
- ✅ `FeedbackType` marshaling/unmarshaling (JSON)
- ✅ `FeedbackType.Name()` method
- ✅ `Sentiment` marshaling/unmarshaling (JSON)
- ✅ `Sentiment.Name()` method
- ✅ `Sentiment.Category()` method
- ✅ Invalid enum value handling
- ✅ Edge cases (empty strings, unknown values)

### Entity Tests (`entity_test.go`)
- ✅ `Feedback` JSON marshaling/unmarshaling
- ✅ `Feedback` with all feedback types (bug, feature, general)
- ✅ `Feedback` with all sentiments (neutral, positive, negative)
- ✅ Optional field handling (`sentimentScore`)
- ✅ Empty tags array handling
- ✅ `Vote` JSON marshaling/unmarshaling
- ✅ `Vote` with authenticated users (`userId`)
- ✅ `Vote` with anonymous users (`ipAddress`)
- ✅ `Vote` with fully anonymous (no userId, no ipAddress)

## Test Patterns

### Table-Driven Tests

Most tests use the table-driven pattern for comprehensive coverage:

```go
tests := []struct {
    name     string
    input    interface{}
    expected interface{}
    wantErr  bool
}{
    {
        name:     "Test case description",
        input:    someInput,
        expected: expectedOutput,
        wantErr:  false,
    },
    // ... more test cases
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // Test logic here
    })
}
```

### Helper Functions

Available in `helpers_test.go`:
- `assertJSONEqual(t, expected, actual)` - Compare JSON strings semantically
- `floatPtr(f)` - Create float64 pointer
- `intPtr(i)` - Create int pointer
- `stringPtr(s)` - Create string pointer

## Writing New Tests

When adding new functionality, follow these guidelines:

1. **Use table-driven tests** for multiple input/output scenarios
2. **Test edge cases** (nil, empty, invalid values)
3. **Use descriptive test names** that explain what is being tested
4. **Use subtests** (`t.Run()`) to organize related test cases
5. **Add helper functions** to `helpers_test.go` when needed
6. **Test both success and failure** paths

### Example Test Template

```go
func TestNewFeature(t *testing.T) {
    tests := []struct {
        name     string
        input    YourType
        expected ExpectedType
        wantErr  bool
    }{
        {
            name:     "Valid case",
            input:    validInput,
            expected: expectedOutput,
            wantErr:  false,
        },
        {
            name:     "Invalid case",
            input:    invalidInput,
            expected: nil,
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := YourFunction(tt.input)

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
```

## Continuous Integration

These tests are designed to run in CI/CD pipelines:

```bash
# In your CI script
go test ./tests/... -v -race -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## Best Practices

1. **Run tests before committing** code
2. **Maintain high test coverage** (aim for >80%)
3. **Test public APIs** thoroughly
4. **Keep tests fast** - avoid external dependencies in unit tests
5. **Use meaningful assertions** with clear error messages
6. **Document complex test scenarios** with comments

## Future Enhancements

Potential additions to the test suite:

- [ ] Integration tests with database
- [ ] Benchmark tests for performance-critical code
- [ ] Fuzz testing for input validation
- [ ] Mock database for repository layer tests
- [ ] E2E tests for full workflows

---

**Note:** The demo file (`demo_models.go`) in the root directory provides a working example of how to use the models. Run it with `make demo` or `go run demo_models.go`.

