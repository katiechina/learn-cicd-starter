package auth

import (
	"net/http"
	"testing"
)

// TestGetAPIKeyWithComplexKey tests with a complex API key containing special characters
func TestGetAPIKeyWithComplexKey(t *testing.T) {
	headers := http.Header{
		"Authorization": []string{"ApiKey abc123!@#$%^&*()_+-=[]{}|;:,.<>?/`~"},
	}
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "abc123!@#$%^&*()_+-=[]{}|;:,.<>?/`~"
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

// TestGetAPIKeyWithMultipleSpaces tests with multiple spaces between ApiKey and the actual key
func TestGetAPIKeyWithMultipleSpaces(t *testing.T) {
	headers := http.Header{
		"Authorization": []string{"ApiKey    my-api-key"},
	}
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "my-api-key"
	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
